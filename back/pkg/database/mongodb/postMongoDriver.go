package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database/status"
)

type PostMongoDriver struct {
	collection *mongo.Collection
}

func (p PostMongoDriver) Insert(post *post.Post) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Microsecond)
	defer cancel()

	if _, err := p.collection.InsertOne(ctx, post); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL
}

func (p PostMongoDriver) Get(name string) (*post.Post, status.QueryStatus) {
	panic("implement me")
}

func (p PostMongoDriver) Delete(name string) status.QueryStatus {
	panic("implement me")
}

func (p PostMongoDriver) Update(name string) status.QueryStatus {
	panic("implement me")
}

func NewPostMongoDriver(db, collection string) *PostMongoDriver {
	return &PostMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
