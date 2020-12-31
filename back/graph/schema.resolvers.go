package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"yes-blog/graph/generated"
	"yes-blog/graph/model"
	controller "yes-blog/internal/controller/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, target model.TargetUser) (*model.User, error) {
	newUser, err := controller.GetUserController().Create(target.Username, target.Password)
	return reformatUser(newUser),err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, name string) (string, error) {
	return name, controller.GetUserController().Delete(&name)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, target model.TargetUser) (*model.User, error) {

	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timeline(ctx context.Context, start int, amount int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	all, err := controller.GetUserController().GetAll(int64(start), int64(amount))
	return reformatUsers(all), err
}

func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	blogUser, err := controller.GetUserController().Get(&name)
	return reformatUser(blogUser), err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
