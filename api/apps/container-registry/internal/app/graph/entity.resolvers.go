package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	httpServer "github.com/kloudlite/api/pkg/http-server"
	"github.com/kloudlite/api/pkg/repos"
)

// FindBuildByID is the resolver for the findBuildByID field.
func (r *entityResolver) FindBuildByID(ctx context.Context, id repos.ID) (*entities.Build, error) {
	sess := httpServer.GetSession[*common.AuthSession](ctx)
	if sess == nil {
		return nil, fiber.ErrUnauthorized
	}
	m := httpServer.GetHttpCookies(ctx)
	klAccount := m[r.Env.AccountCookieName]
	if klAccount == "" {
		return nil, errors.Newf("no cookie named %q present in request", r.Env.AccountCookieName)
	}

	nctx := context.WithValue(ctx, "user-session", sess)
	nctx = context.WithValue(nctx, "account-name", klAccount)

	cc, err := toRegistryContext(nctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetBuild(cc, id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }