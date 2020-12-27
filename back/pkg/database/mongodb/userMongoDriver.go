package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database/status"
)

type UserMongoDriver struct {
	collection *mongo.Collection
}

func (u UserMongoDriver) GetAll(amount int64) ([]*user.User, status.QueryStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	var result []*user.User
	if cur, err := u.collection.Find(ctx, bson.D{}); err != nil {
		return nil, status.FAILED
	} else {
		defer cur.Close(ctx)
		for cur.Next(context.Background()) {
			if amount >0{
				break
			}
			amount--
			var user user.User
			_ = cur.Decode(&user)
			result = append(result, &user)
		}
		return result,status.SUCCESSFUL
	}
}

func (u UserMongoDriver) Insert(user *user.User) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL

}

func (u UserMongoDriver) Get(name string) (*user.User, status.QueryStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	var result user.User
	if err := u.collection.FindOne(ctx, fmt.Sprintf("{name:%s}", name)).Decode(&result); err != nil {
		return &result, status.FAILED
	}
	return &result, status.SUCCESSFUL
}

func (u UserMongoDriver) Delete(name string) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	if _, err := u.collection.DeleteOne(ctx, fmt.Sprintf("{name:%s}", name)); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL
}

func (u UserMongoDriver) Update(user *user.User) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	if _, err := u.collection.UpdateOne(ctx, fmt.Sprintf("{name:%s}", user.Name), user); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL

}

func NewUserMongoDriver(db, collection string) *UserMongoDriver {
	return &UserMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
