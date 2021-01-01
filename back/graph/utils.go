package graph

import (
	"context"
	"yes-blog/graph/model"
	"yes-blog/internal/middleware/auth"
	"yes-blog/internal/middleware/ggcontext"
	"yes-blog/internal/model/post"
	"yes-blog/internal/model/user"
)

/* some useful functions to convert model objects from our models to graphql models
 */
func reformatUsers(all []*user.User) []*model.User {
	var result []*model.User
	for _, blogUser := range all {
		result = append(result, reformatUser(blogUser))
	}
	return result
}

func reformatUser(blogUser *user.User) *model.User {
	var graphUser = &model.User{
		ID:   blogUser.ID.Hex(),
		Name: blogUser.Name,
		Email: blogUser.Email,
	}
	graphUser.Posts = reformatPosts(blogUser.Posts, graphUser)
	return graphUser
}

func reformatPost(blogPost *post.Post, graphUser *model.User) *model.Post {
	return &model.Post{
		ID:        blogPost.ID.Hex(),
		CreatedBy: graphUser,
		CreatedAt: int(blogPost.TimeStamp),
		Content:   blogPost.Body,
		Title:     blogPost.Title,
	}
}

func reformatPosts(blogPosts []*post.Post, graphUser *model.User) []*model.Post {
	var posts []*model.Post
	for _, p := range blogPosts {
		posts = append(posts, reformatPost(p, graphUser))
	}
	return posts
}

func reformatAllPosts(blogPosts []*post.Post, graphUser *model.User) []*model.Post {
	var posts []*model.Post
	for _, p := range blogPosts {
		posts = append(posts, reformatPost(p, graphUser))
	}
	return posts
}

func reformatAllSeparatePosts(blogPosts []*post.Post, graphUser []*model.User) []*model.Post {
	var posts []*model.Post
	for i, _ := range blogPosts {
		posts = append(posts, reformatPost(blogPosts[i], graphUser[i]))
	}
	return posts
}


func extractUsernameFromContext(ctx context.Context) string {
	ginContext, _ := ggcontext.GinContextFromContext(ctx)
	return auth.ForContext(ginContext)
}