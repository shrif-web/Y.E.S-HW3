package database

import (
	"yes-blog/internal/model/post"
	"yes-blog/internal/model/user"
)

type PostDBDriver interface {
	Insert(post *post.Post) (*user.User, error)
	Get(postID string) (*post.Post, *user.User, error)
	GetAll(startIndex, amount int) ([]*post.Post, []*user.User, error)
	GetByUser(userID string) ([]*post.Post, *user.User, error)
	Delete(postID string, authorName string) error
	Update(post *post.Post) error
}
