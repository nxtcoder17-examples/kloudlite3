package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/comms/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/comms/internal/app/graph/model"
	"github.com/kloudlite/api/apps/comms/internal/domain/entities"
	fn "github.com/kloudlite/api/pkg/functions"
)

// CreationTime is the resolver for the creationTime field.
func (r *notificationConfResolver) CreationTime(ctx context.Context, obj *entities.NotificationConf) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("obj is nil")
	}

	return obj.CreationTime.Format(time.RFC3339), nil
}

// Email is the resolver for the email field.
func (r *notificationConfResolver) Email(ctx context.Context, obj *entities.NotificationConf) (*model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesEmail, error) {
	if obj == nil {
		return nil, fmt.Errorf("obj is nil")
	}

	return fn.JsonConvertP[model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesEmail](obj.Email)
}

// Slack is the resolver for the slack field.
func (r *notificationConfResolver) Slack(ctx context.Context, obj *entities.NotificationConf) (*model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesSlack, error) {
	if obj == nil {
		return nil, fmt.Errorf("obj is nil")
	}

	return fn.JsonConvertP[model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesSlack](obj.Slack)
}

// Telegram is the resolver for the telegram field.
func (r *notificationConfResolver) Telegram(ctx context.Context, obj *entities.NotificationConf) (*model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesTelegram, error) {
	if obj == nil {
		return nil, fmt.Errorf("obj is nil")
	}

	return fn.JsonConvertP[model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesTelegram](obj.Telegram)
}

// UpdateTime is the resolver for the updateTime field.
func (r *notificationConfResolver) UpdateTime(ctx context.Context, obj *entities.NotificationConf) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("obj is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Webhook is the resolver for the webhook field.
func (r *notificationConfResolver) Webhook(ctx context.Context, obj *entities.NotificationConf) (*model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesWebhook, error) {
	if obj.Webhook == nil {
		return nil, nil
	}

	return fn.JsonConvertP[model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesWebhook](obj.Webhook)
}

// Email is the resolver for the email field.
func (r *notificationConfInResolver) Email(ctx context.Context, obj *entities.NotificationConf, data *model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesEmailIn) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}

	gckacideeci, err := fn.JsonConvertP[entities.Email](data)
	if err != nil {
		return err
	}

	obj.Email = gckacideeci

	return nil
}

// Slack is the resolver for the slack field.
func (r *notificationConfInResolver) Slack(ctx context.Context, obj *entities.NotificationConf, data *model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesSlackIn) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}

	gckacideesci, err := fn.JsonConvertP[entities.Slack](data)
	if err != nil {
		return err
	}

	obj.Slack = gckacideesci
	return nil
}

// Telegram is the resolver for the telegram field.
func (r *notificationConfInResolver) Telegram(ctx context.Context, obj *entities.NotificationConf, data *model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesTelegramIn) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}

	gckacideetci, err := fn.JsonConvertP[entities.Telegram](data)
	if err != nil {
		return err
	}

	obj.Telegram = gckacideetci

	return nil
}

// Webhook is the resolver for the webhook field.
func (r *notificationConfInResolver) Webhook(ctx context.Context, obj *entities.NotificationConf, data *model.GithubComKloudliteAPIAppsCommsInternalDomainEntitiesWebhookIn) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}

	gckacideesci, err := fn.JsonConvertP[entities.Webhook](data)
	if err != nil {
		return err
	}

	obj.Webhook = gckacideesci
	return nil
}

// NotificationConf returns generated.NotificationConfResolver implementation.
func (r *Resolver) NotificationConf() generated.NotificationConfResolver {
	return &notificationConfResolver{r}
}

// NotificationConfIn returns generated.NotificationConfInResolver implementation.
func (r *Resolver) NotificationConfIn() generated.NotificationConfInResolver {
	return &notificationConfInResolver{r}
}

type notificationConfResolver struct{ *Resolver }
type notificationConfInResolver struct{ *Resolver }
