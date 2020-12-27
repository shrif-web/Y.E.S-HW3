package controller

import (
	"errors"
	"yes-blog/graph/model"
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database/status"
)

func (c *UserController) GetAll(start,amount int64) ([]*model.User, error) {
	all, err := c.dbDriver.GetAll(start,amount)
	if err == status.FAILED {
		return nil, errors.New("couldn't fetch the users required")
	}

	return reformatUsers(all), nil
}

func reformatUsers(all []*user.User) []*model.User {
	var result []*model.User
	for _, blogUser := range all {
		var graphUser = &model.User{
			ID:   blogUser.ID,
			Name: blogUser.Name,
		}
		graphUser.Posts = reformatPosts(blogUser, graphUser)
		result = append(result, graphUser)
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