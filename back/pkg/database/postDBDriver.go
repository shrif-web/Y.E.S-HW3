package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-blog/internal/model/post"
	"yes-blog/internal/model/user"
)

type PostDBDriver interface {
	Insert(post *post.Post, authorName string) error
	Get(postID primitive.ObjectID) (*post.Post, *user.User, error)
	GetAll(startIndex, amount int) ([]*post.Post, []*user.User, error)
	GetByUser(userName string) ([]*post.Post, *user.User, error)
	Delete(postID primitive.ObjectID, authorName string) error
	Update(post *post.Post) error
}
