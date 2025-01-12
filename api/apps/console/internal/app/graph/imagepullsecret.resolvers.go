package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"time"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/entities"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *imagePullSecretResolver) CreationTime(ctx context.Context, obj *entities.ImagePullSecret) (string, error) {
	if obj == nil {
		return "", errNilImagePullSecret
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// Format is the resolver for the format field.
func (r *imagePullSecretResolver) Format(ctx context.Context, obj *entities.ImagePullSecret) (model.GithubComKloudliteAPIAppsConsoleInternalEntitiesPullSecretFormat, error) {
	if obj == nil {
		return "", errNilImagePullSecret
	}
	return model.GithubComKloudliteAPIAppsConsoleInternalEntitiesPullSecretFormat(obj.Format), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *imagePullSecretResolver) UpdateTime(ctx context.Context, obj *entities.ImagePullSecret) (string, error) {
	if obj == nil {
		return "", errNilImagePullSecret
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Format is the resolver for the format field.
func (r *imagePullSecretInResolver) Format(ctx context.Context, obj *entities.ImagePullSecret, data model.GithubComKloudliteAPIAppsConsoleInternalEntitiesPullSecretFormat) error {
	if obj == nil {
		return errNilImagePullSecret
	}

	obj.Format = entities.PullSecretFormat(data)
	return nil
}

// Metadata is the resolver for the metadata field.
func (r *imagePullSecretInResolver) Metadata(ctx context.Context, obj *entities.ImagePullSecret, data *v1.ObjectMeta) error {
	if obj == nil {
		return errNilImagePullSecret
	}

	if data != nil {
		obj.ObjectMeta = *data
	}

	return nil
}

// ImagePullSecret returns generated.ImagePullSecretResolver implementation.
func (r *Resolver) ImagePullSecret() generated.ImagePullSecretResolver {
	return &imagePullSecretResolver{r}
}

// ImagePullSecretIn returns generated.ImagePullSecretInResolver implementation.
func (r *Resolver) ImagePullSecretIn() generated.ImagePullSecretInResolver {
	return &imagePullSecretInResolver{r}
}

type imagePullSecretResolver struct{ *Resolver }
type imagePullSecretInResolver struct{ *Resolver }
