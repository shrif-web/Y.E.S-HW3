package database

import (
	"yes-blog/internal/model/post"
)

type PostDBDriver interface {
	Insert(post *post.Post) error
	Get(postID string) (*post.Post, error)
	GetAll(startID string, amount int) ([]*post.Post, error)
	GetByUser(userID string) ([]*post.Post, error)
	Delete(postID string, authorName string) error
	Update(post *post.Post) error
}
