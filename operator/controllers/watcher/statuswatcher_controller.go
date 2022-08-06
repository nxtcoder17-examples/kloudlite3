package watcher

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	types2 "k8s.io/apimachinery/pkg/types"
	crdsv1 "operators.kloudlite.io/apis/crds/v1"
	serverlessv1 "operators.kloudlite.io/apis/serverless/v1"
	"operators.kloudlite.io/env"
	"operators.kloudlite.io/lib/constants"
	fn "operators.kloudlite.io/lib/functions"
	"operators.kloudlite.io/lib/logging"
	rApi "operators.kloudlite.io/lib/operator"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// StatusWatcherReconciler reconciles a StatusWatcher object
type StatusWatcherReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Env    *env.Env
	*Notifier
	Logger logging.Logger
}

func (r *StatusWatcherReconciler) GetName() string {
	return "status-watcher"
}

func parseGroup(b64GroupName string) (*schema.GroupVersionKind, error) {
	gName, err := base64.StdEncoding.DecodeString(b64GroupName)
	if err != nil {
		return nil, err
	}
	var gvk schema.GroupVersionKind
	s := strings.Split(string(gName), ", ")
	gv := strings.Split(s[0], "/")
	gvk.Group = strings.TrimSpace(gv[0])
	gvk.Version = strings.TrimSpace(gv[1])
	if _, err := fmt.Sscanf(s[1], "Kind=%s", &gvk.Kind); err != nil {
		return nil, err
	}
	return &gvk, nil
}

func (r *StatusWatcherReconciler) SendStatusEvents(ctx context.Context, obj client.Object) (ctrl.Result, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return ctrl.Result{}, nil
	}

	var j struct {
		Status rApi.Status `json:"status"`
	}

	if err := json.Unmarshal(b, &j); err != nil {
		return ctrl.Result{}, nil
	}

	klMetadata := ExtractMetadata(obj)

	if obj.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(obj, constants.StatusWatcherFinalizer) {
			if err := r.notify(ctx, getMsgKey(obj), klMetadata, j.Status, Stages.Deleted); err != nil {
				return ctrl.Result{}, err
			}
			return r.RemoveWatcherFinalizer(ctx, obj)
		}
		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(obj, constants.StatusWatcherFinalizer) {
		return r.AddWatcherFinalizer(ctx, obj)
	}
	if err := r.notify(ctx, getMsgKey(obj), klMetadata, j.Status, Stages.Exists); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=watcher.kloudlite.io,resources=statuswatchers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=watcher.kloudlite.io,resources=statuswatchers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=watcher.kloudlite.io,resources=statuswatchers/finalizers,verbs=update

func (r *StatusWatcherReconciler) Reconcile(ctx context.Context, oReq ctrl.Request) (ctrl.Result, error) {
	var wName WrappedName
	if err := json.Unmarshal([]byte(oReq.Name), &wName); err != nil {
		return ctrl.Result{}, nil
	}

	gvk, err := parseGroup(wName.Group)
	if err != nil {
		r.Logger.Errorf(err, "badly formatted group-version-kind (%s) received, aborting ...", wName.Group)
		return ctrl.Result{}, nil
	}

	if gvk == nil {
		return ctrl.Result{}, nil
	}

	logger := r.Logger.WithName(fn.NN(oReq.Namespace, wName.Name).String()).WithKV("RefKind", gvk.String())
	logger.Infof("request received ...")

	tm := metav1.TypeMeta{Kind: gvk.Kind, APIVersion: fmt.Sprintf("%s/%s", gvk.Group, gvk.Version)}
	obj, err := rApi.Get(ctx, r.Client, fn.NN(oReq.Namespace, wName.Name), fn.NewUnstructured(tm))
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	return r.SendStatusEvents(ctx, obj)
}

func (r *StatusWatcherReconciler) AddWatcherFinalizer(ctx context.Context, obj client.Object) (ctrl.Result, error) {
	controllerutil.AddFinalizer(obj, constants.StatusWatcherFinalizer)
	return ctrl.Result{}, r.Update(ctx, obj)
}

func (r *StatusWatcherReconciler) RemoveWatcherFinalizer(ctx context.Context, obj client.Object) (ctrl.Result, error) {
	controllerutil.RemoveFinalizer(obj, constants.StatusWatcherFinalizer)
	return ctrl.Result{}, r.Update(ctx, obj)
}

// SetupWithManager sets up the controller with the Manager.
func (r *StatusWatcherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.Logger = r.Logger.WithName("status-watcher")

	builder := ctrl.NewControllerManagedBy(mgr)
	builder.For(&crdsv1.Project{})

	watchList := []client.Object{
		&crdsv1.Project{},
		&crdsv1.App{},
		&serverlessv1.Lambda{},
		&crdsv1.ManagedService{},
		&crdsv1.ManagedResource{},
		&crdsv1.Router{},
	}

	for _, object := range watchList {
		builder.Watches(
			&source.Kind{Type: object},
			handler.EnqueueRequestsFromMapFunc(
				func(obj client.Object) []reconcile.Request {
					b64Group := base64.StdEncoding.EncodeToString(
						[]byte(obj.GetAnnotations()[constants.AnnotationKeys.GroupVersionKind]),
					)

					if len(b64Group) == 0 {
						return nil
					}

					wName, err := WrappedName{Name: obj.GetName(), Group: b64Group}.String()
					if err != nil {
						return nil
					}
					return []reconcile.Request{
						{
							NamespacedName: types2.NamespacedName{Namespace: obj.GetNamespace(), Name: wName},
						},
					}
				},
			),
		)
	}
	return builder.Complete(r)
}
