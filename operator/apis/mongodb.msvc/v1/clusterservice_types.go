package v1

import (
	ct "github.com/kloudlite/operator/apis/common-types"
	"github.com/kloudlite/operator/pkg/constants"
	rApi "github.com/kloudlite/operator/pkg/operator"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterServiceSpec defines the desired state of ClusterService
type ClusterServiceSpec struct {
	Resources        ct.Resources `json:"resources"`
	OutputSecretName *string      `json:"outputSecretName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterService is the Schema for the clusterservices API
type ClusterService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterServiceSpec `json:"spec"`
	Status rApi.Status        `json:"status,omitempty"`
}

func (cs *ClusterService) EnsureGVK() {
	if cs != nil {
		cs.SetGroupVersionKind(GroupVersion.WithKind("ClusterService"))
	}
}

func (c *ClusterService) GetStatus() *rApi.Status {
	return &c.Status
}

func (c *ClusterService) GetEnsuredLabels() map[string]string {
	return map[string]string{
		constants.MsvcNameKey: c.Name,
	}
}

func (c *ClusterService) GetEnsuredAnnotations() map[string]string {
	return map[string]string{
		constants.AnnotationKeys.GroupVersionKind: GroupVersion.WithKind("ClusterService").String(),
	}
}

// +kubebuilder:object:root=true

// ClusterServiceList contains a list of ClusterService
type ClusterServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterService{}, &ClusterServiceList{})
}
