package database

import "yes-blog/internal/model/post"


type PostDBDriver interface {
	Insert() (*post.Post,QueryStatus)
	Get(name string) (*post.Post,QueryStatus)
	Delete(name string) QueryStatus
	Update(name string)	QueryStatus
}