package database

import (
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database/status"
)

type PostDBDriver interface {
	Insert(post *post.Post) status.QueryStatus
	Get(name string) (*post.Post,status.QueryStatus)
	Delete(name string) status.QueryStatus
	Update(name string)	status.QueryStatus
}