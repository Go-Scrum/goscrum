package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"goscrum/goscrum/server/graph/gqlgen_gen"
	"goscrum/goscrum/server/graph/model"
)

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int, search *string) (*model.UserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SlackUsers(ctx context.Context, limit *int, offset *int, search *string) (*model.SlackUserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlgen_gen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen_gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
