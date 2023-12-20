package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *projectResolver) CreationTime(ctx context.Context, obj *entities.Project) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *projectResolver) ID(ctx context.Context, obj *entities.Project) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *projectResolver) Spec(ctx context.Context, obj *entities.Project) (*model.GithubComKloudliteOperatorApisCrdsV1ProjectSpec, error) {
	m := &model.GithubComKloudliteOperatorApisCrdsV1ProjectSpec{}
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *projectResolver) UpdateTime(ctx context.Context, obj *entities.Project) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *projectInResolver) Metadata(ctx context.Context, obj *entities.Project, data *v1.ObjectMeta) error {
	obj.ObjectMeta = *data
	return nil
}

// Spec is the resolver for the spec field.
func (r *projectInResolver) Spec(ctx context.Context, obj *entities.Project, data *model.GithubComKloudliteOperatorApisCrdsV1ProjectSpecIn) error {
	if obj == nil {
		return errors.Newf("resource is nil")
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// ProjectIn returns generated.ProjectInResolver implementation.
func (r *Resolver) ProjectIn() generated.ProjectInResolver { return &projectInResolver{r} }

type projectResolver struct{ *Resolver }
type projectInResolver struct{ *Resolver }
