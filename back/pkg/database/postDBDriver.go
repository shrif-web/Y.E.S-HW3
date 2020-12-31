package database

import (
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database/status"
)

type PostDBDriver interface {
	Insert(post *post.Post) status.QueryStatus
	Get(postID string) (*post.Post, status.QueryStatus)
	GetAll(startID, amount uint64) ([]*post.Post, status.QueryStatus)
	GetByUser(userID string) ([]*post.Post, status.QueryStatus)
	Delete(postID string) status.QueryStatus
	Update(post *post.Post) (status.QueryStatus, error)
}
