package graph

import (
	"yes-blog/graph/model"
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

func reformatPosts(blogUser *user.User, graphUser *model.User) []*model.Post {
	var posts []*model.Post
	for _, p := range blogUser.Posts {
		posts = append(posts, &model.Post{
			ID:        p.ID,
			Auther:    graphUser,
			Timestamp: int(p.TimeStamp),
			Body:      p.Body,
			Title:     p.Title,
		})
	}
	return posts
}
func reformatUser(blogUser *user.User) *model.User {
	var graphUser = &model.User{
		ID:   blogUser.ID,
		Name: blogUser.Name,
	}
	graphUser.Posts = reformatPosts(blogUser, graphUser)
	return graphUser
}