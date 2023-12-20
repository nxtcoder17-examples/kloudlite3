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
func (r *managedResourceResolver) CreationTime(ctx context.Context, obj *entities.ManagedResource) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *managedResourceResolver) ID(ctx context.Context, obj *entities.ManagedResource) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *managedResourceResolver) Spec(ctx context.Context, obj *entities.ManagedResource) (*model.GithubComKloudliteOperatorApisCrdsV1ManagedResourceSpec, error) {
	m := &model.GithubComKloudliteOperatorApisCrdsV1ManagedResourceSpec{}
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *managedResourceResolver) UpdateTime(ctx context.Context, obj *entities.ManagedResource) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return obj.BaseEntity.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *managedResourceInResolver) Metadata(ctx context.Context, obj *entities.ManagedResource, data *v1.ObjectMeta) error {
	obj.ObjectMeta = *data
	return nil
}

// Spec is the resolver for the spec field.
func (r *managedResourceInResolver) Spec(ctx context.Context, obj *entities.ManagedResource, data *model.GithubComKloudliteOperatorApisCrdsV1ManagedResourceSpecIn) error {
	if obj == nil {
		return errors.Newf("resource is nil")
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// ManagedResource returns generated.ManagedResourceResolver implementation.
func (r *Resolver) ManagedResource() generated.ManagedResourceResolver {
	return &managedResourceResolver{r}
}

// ManagedResourceIn returns generated.ManagedResourceInResolver implementation.
func (r *Resolver) ManagedResourceIn() generated.ManagedResourceInResolver {
	return &managedResourceInResolver{r}
}

type managedResourceResolver struct{ *Resolver }
type managedResourceInResolver struct{ *Resolver }
