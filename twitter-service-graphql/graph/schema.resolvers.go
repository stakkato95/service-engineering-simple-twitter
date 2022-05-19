package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/graph/generated"
	"github.com/stakkato95/twitter-service-graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	logger.Info("not implemented")
	return "hello", nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	logger.Info("not implemented")
	return "world", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
