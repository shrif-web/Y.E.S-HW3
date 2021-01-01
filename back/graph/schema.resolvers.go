package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"yes-blog/graph/generated"
	"yes-blog/graph/model"
	postController "yes-blog/internal/controller/post"
	userController "yes-blog/internal/controller/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, target model.TargetUser) (model.CreateUserPayload, error) {
	newUser, err := userController.GetUserController().Create(target.Username, target.Password)
	if err != nil {
		switch err.(type) {
		case model.DuplicateUsernameException:
			return err.(model.DuplicateUsernameException), nil
		case model.InternalServerException:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatUser(newUser), err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, name string) (string, error) {
	return name, userController.GetUserController().Delete(&name)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, target string, toBe model.ToBeUser) (model.UpdateUserPayload, error) {
	update, err := userController.GetUserController().Update(target, toBe)
	if err != nil {
		switch err.(type) {
		case model.NoUserFoundException:
			return err.(model.NoUserFoundException), nil
		case model.InternalServerException:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatUser(update), err
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (model.LoginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.TargetPost) (*model.Post, error) {
	newPost, err := postController.GetPostController().CreatePost(input.Title, input.Content, input.AuthorName)
	if err != nil {
		return nil, err
	}
	return reformatPost(newPost, &model.User{
		ID:    "",
		Name:  input.AuthorName,
		Posts: nil,
	}), nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, targetID string, authorName string) (string, error) {
	message, err := postController.GetPostController().DeletePost(targetID, authorName)
	if err != nil {
		return fmt.Sprint(err), err
	}
	return message, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, targetID string, input model.TargetPost) (string, error) {
	message, err := postController.GetPostController().UpdatePost(targetID, input.Title, input.Content, input.AuthorName)
	if err != nil {
		return fmt.Sprint(err), err
	}
	return message, nil
}

func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	blogUser, err := userController.GetUserController().Get(&name)
	return reformatUser(blogUser), err
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	all, err := userController.GetUserController().GetAll(int64(start), int64(amount))
	return reformatUsers(all), err
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	newPost, err := postController.GetPostController().GetPost(id)
	if err != nil {
		return nil, err
	}
	return reformatPost(newPost, &model.User{
		ID:    "",
		Name:  newPost.Author,
		Posts: nil,
	}), nil
}

func (r *queryResolver) Posts(ctx context.Context, start int, amount int) ([]*model.Post, error) {
	posts, err := postController.GetPostController().GetAllPosts(strconv.Itoa(start), amount)
	if err != nil {
		return nil, err
	}
	return reformatAllPosts(posts), nil
}

func (r *queryResolver) PostsOfUser(ctx context.Context, userName string) ([]*model.Post, error) {
	posts, err := postController.GetPostController().GetPostByUser(userName)
	if err != nil {
		return nil, err
	}
	return reformatPosts(posts, &model.User{
		ID:    "",
		Name:  userName,
		Posts: nil,
	}), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
