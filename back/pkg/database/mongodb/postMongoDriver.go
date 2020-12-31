package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database/status"
)

type PostMongoDriver struct {
	collection *mongo.Collection
}

func (p PostMongoDriver) Insert(post *post.Post) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	if _, err := p.collection.InsertOne(ctx, post); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL
}

func (p PostMongoDriver) Get(postID string) (*post.Post, status.QueryStatus) {
	panic("implement me")
}

func (p PostMongoDriver) GetAll(startID, amount uint64) ([]*post.Post, status.QueryStatus) {
	panic("implement me")
}

func (p PostMongoDriver) GetByUser(userID string) ([]*post.Post, status.QueryStatus) {
	panic("implement me")
}

func (p PostMongoDriver) Delete(postID string) status.QueryStatus {
	panic("implement me")
}

func (p PostMongoDriver) Update(pst *post.Post) (status.QueryStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	var result post.Post
	err := p.collection.FindOne(ctx, fmt.Sprintf("{id:%s authorID:%s}", pst.GetID(), pst.GetAuthorID())).Decode(&result)
	if err != nil {
		return status.FAILED, errors.New("cannot find the post " + pst.GetID())
	}
	if result.GetAuthorID() != pst.GetAuthorID() {
		return status.FAILED, errors.New("user " + result.GetAuthorID() + " not allowed to edit post " + pst.GetID())
	}
	_, errr := p.collection.UpdateOne(ctx, fmt.Sprintf("{title:%s body:%s timeStamp:%d}", pst.GetTitle(), pst.GetBody(), pst.GetTimeStamp()), pst)
	if errr != nil {
		return status.FAILED, errr
	}
	return status.SUCCESSFUL, nil
}

func NewPostMongoDriver(db, collection string) *PostMongoDriver {
	return &PostMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
