package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	generated1 "kloudlite.io/apps/container-registry/internal/app/graph/generated"
	"kloudlite.io/apps/container-registry/internal/app/graph/model"
	"kloudlite.io/apps/container-registry/internal/domain/entities"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/repos"
)

// CrCreateRepo is the resolver for the cr_createRepo field.
func (r *mutationResolver) CrCreateRepo(ctx context.Context, repository entities.Repository) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.Domain.CreateRepository(cc, repository.Name); err != nil {
		return false, err
	}

	return true, nil
}

// CrCreateCred is the resolver for the cr_createCred field.
func (r *mutationResolver) CrCreateCred(ctx context.Context, credential entities.Credential) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.Domain.CreateCredential(cc, credential); err != nil {
		return false, err
	}

	return true, nil
}

// CrDeleteRepo is the resolver for the cr_deleteRepo field.
func (r *mutationResolver) CrDeleteRepo(ctx context.Context, name string) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteRepository(cc, name); err != nil {
		return false, err
	}

	return true, nil
}

// CrDeleteCred is the resolver for the cr_deleteCred field.
func (r *mutationResolver) CrDeleteCred(ctx context.Context, name string, username string) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteCredential(cc, name, username); err != nil {
		return false, err
	}
	return true, nil
}

// CrDeleteTag is the resolver for the cr_deleteTag field.
func (r *mutationResolver) CrDeleteTag(ctx context.Context, repoName string, tagName string) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteRepositoryTag(cc, repoName, tagName); err != nil {
		return false, err
	}
	return true, nil
}

// CrListRepos is the resolver for the cr_listRepos field.
func (r *queryResolver) CrListRepos(ctx context.Context, search *model.SearchRepos, pagination *repos.CursorPagination) (*model.RepositoryPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListRepositories(cc, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.RepositoryEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.RepositoryEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.RepositoryPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: len(records),
	}

	return m, nil
}

// CrListCreds is the resolver for the cr_listCreds field.
func (r *queryResolver) CrListCreds(ctx context.Context, search *model.SearchCreds, pagination *repos.CursorPagination) (*model.CredentialPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListCredentials(cc, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.CredentialEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.CredentialEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.CredentialPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: len(records),
	}

	return m, nil
}

// CrListTags is the resolver for the cr_listTags field.
func (r *queryResolver) CrListTags(ctx context.Context, repoName string, search *model.SearchRepos, pagination *repos.CursorPagination) (*model.TagPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListRepositoryTags(cc, repoName, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.TagEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.TagEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}
	m := &model.TagPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: len(records),
	}

	return m, nil
}

// CrGetCredToken is the resolver for the cr_getCredToken field.
func (r *queryResolver) CrGetCredToken(ctx context.Context, username string) (string, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return "", err
	}

	token, err := r.Domain.GetToken(cc, username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
