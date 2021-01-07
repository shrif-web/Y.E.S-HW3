package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"yes-blog/graph/generated"
	"yes-blog/graph/model"
	postController "yes-blog/internal/controller/post"
	userController "yes-blog/internal/controller/user"
	"yes-blog/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, target model.TargetUser) (model.CreateUserPayload, error) {
	println("user: " + extractUsernameFromContext(ctx))
	newUser, err := userController.GetUserController().Create(target.Username, target.Password, target.Email)
	if err != nil {
		switch err.(type) {
		case model.DuplicateUsernameException:
			return err.(model.DuplicateUsernameException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatUser(newUser), err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, name string) (model.DeletePostPayload, error) {
	err := userController.GetUserController().Delete(extractUsernameFromContext(ctx), name)
	if err != nil {
		switch err.(type) {
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.NoUserFoundException:
			return err.(model.NoUserFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return model.OperationSuccessfull{}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, toBe model.ToBeUser) (model.UpdateUserPayload, error) {
	update, err := userController.GetUserController().Update(extractUsernameFromContext(ctx), toBe)
	if err != nil {
		switch err.(type) {
		case model.NoUserFoundException:
			return err.(model.NoUserFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatUser(update), err
}

func (r *mutationResolver) Promote(ctx context.Context, target string) (model.AdminPayload, error) {
	if err := userController.GetUserController().Promote(extractUsernameFromContext(ctx), target); err != nil {
		switch err.(type) {
		case model.NoUserFoundException:
			return err.(model.NoUserFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return model.OperationSuccessfull{}, nil
}

func (r *mutationResolver) Demote(ctx context.Context, target string) (model.AdminPayload, error) {
	if err := userController.GetUserController().Demote(extractUsernameFromContext(ctx), target); err != nil {
		switch err.(type) {
		case model.NoUserFoundException:
			return err.(model.NoUserFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return model.OperationSuccessfull{}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (model.LoginPayload, error) {
	token, err := userController.GetUserController().Login(input.Username, input.Password)
	if err != nil {
		switch err.(type) {
		case model.UserPassMissMatchException:
			return err.(model.UserPassMissMatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return model.Token{Token: token}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (model.LoginPayload, error) {
	username := extractUsernameFromContext(ctx)
	if username == "" {
		return model.InternalServerException{Message: "user not found!"}, nil
	}
	// generate new token
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return model.InternalServerException{}, nil
	}
	return model.Token{Token: token}, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.TargetPost) (model.CreatePostPayload, error) {
	newPost, usr, err := postController.GetPostController().CreatePost(input.Title, input.Content, extractUsernameFromContext(ctx))
	if err != nil {
		switch err.(type) {
		case *model.PostEmptyException:
			return err.(*model.PostEmptyException), nil
		case *model.NoUserFoundException:
			return err.(*model.NoUserFoundException), nil
		default:
			return err.(*model.InternalServerException), nil
		}
	}
	return reformatPost(newPost, reformatUser(usr)), nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, targetID string, userName *string) (model.DeletePostPayload, error) {
	message, err := postController.GetPostController().DeletePost(targetID, extractUsernameFromContext(ctx), getUserName(ctx, userName))
	if err != nil {
		switch err.(type) {
		case *model.UserNotAllowedException:
			return err.(*model.UserNotAllowedException), nil
		case *model.NoUserFoundException:
			return err.(*model.NoUserFoundException), nil
		case *model.PostNotFoundException:
			return err.(*model.PostNotFoundException), nil
		default:
			return err.(*model.InternalServerException), nil
		}
	}
	return &model.OperationSuccessfull{Message: message}, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, targetID string, input model.TargetPost, userName *string) (model.UpdatePostPayload, error) {
	message, err := postController.GetPostController().UpdatePost(targetID, input.Title, input.Content, extractUsernameFromContext(ctx), getUserName(ctx, userName))
	if err != nil {
		switch err.(type) {
		case *model.UserNotAllowedException:
			return err.(*model.UserNotAllowedException), nil
		case *model.NoUserFoundException:
			return err.(*model.NoUserFoundException), nil
		case *model.PostNotFoundException:
			return err.(*model.PostNotFoundException), nil
		case *model.PostEmptyException:
			return err.(*model.PostEmptyException), nil
		default:
			return err.(*model.InternalServerException), nil
		}
	}
	return &model.OperationSuccessfull{Message: message}, nil
}

func (r *queryResolver) User(ctx context.Context, name *string) (*model.User, error) {
	//todo error handling
	username := getUserName(ctx, name)
	blogUser, err := userController.GetUserController().Get(&username)
	return reformatUser(blogUser), err
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	all, err := userController.GetUserController().GetAll(int64(start), int64(amount))
	return reformatUsers(all), err
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	newPost, usr, err := postController.GetPostController().GetPost(id)
	if err != nil {
		return nil, err
	}
	return reformatPost(newPost, reformatUser(usr)), nil
}

func (r *queryResolver) Timeline(ctx context.Context, start int, amount int) ([]*model.Post, error) {
	posts, usrs, err := postController.GetPostController().GetAllPosts(start, amount)
	if err != nil {
		return nil, err
	}
	return reformatAllSeparatePosts(posts, reformatUsers(usrs)), nil
}

func (r *queryResolver) PostsOfUser(ctx context.Context, userName *string) ([]*model.Post, error) {
	posts, usr, err := postController.GetPostController().GetPostByUser(getUserName(ctx, userName))
	if err != nil {
		return nil, err
	}
	return reformatAllPosts(posts, reformatUser(usr)), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
