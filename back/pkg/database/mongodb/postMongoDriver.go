package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database"
)

type PostMongoDriver struct {
	client mongo.Client
	db     string
	collection string
}

func NewPostMongoDriver(db,collection string) *PostMongoDriver {
	return &PostMongoDriver{
		client: GetMongoClient(),
		db:     db,
		collection: collection,
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
