package graph

import (
	"yes-blog/graph/model"
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
		ID:   blogUser.ID,
		Name: blogUser.Name,
	}
	graphUser.Posts = reformatPosts(blogUser.Posts, graphUser)
	return graphUser
}

func reformatPost(blogPost *post.Post, graphUser *model.User) *model.Post {
	return &model.Post{
		ID:        blogPost.ID,
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

func reformatAllPosts(blogPosts []*post.Post) []*model.Post {
	var posts []*model.Post
	for _, p := range blogPosts {
		posts = append(posts, reformatPost(p, &model.User{
			ID:    "",
			Name:  p.Author,
			Posts: nil,
		}))
	}
	return posts
}
