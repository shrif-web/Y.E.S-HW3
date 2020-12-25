package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database"
)

type PostMongoDriver struct{
	client mongo.Client
}

func NewPostMongoDriver() *PostMongoDriver {
	return &PostMongoDriver{
		client: GetMongoClient(),
	}
}

func (p PostMongoDriver) Insert() (*post.Post, database.QueryStatus) {
	panic("implement me")
}

func (p PostMongoDriver) Get(name string) (*post.Post, database.QueryStatus) {
	panic("implement me")
}

func (p PostMongoDriver) Delete(name string) database.QueryStatus {
	panic("implement me")
}

func (p PostMongoDriver) Update(name string) database.QueryStatus {
	panic("implement me")
}

