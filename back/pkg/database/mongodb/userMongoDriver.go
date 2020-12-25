package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database"
)

type UserMongoDriver struct {
	client mongo.Client
	db     string
}

func NewUserMongoDriver(db string) *UserMongoDriver {
	return &UserMongoDriver{
		client: GetMongoClient(),
		db:     db,
	}
}

func (u UserMongoDriver) Insert() (*user.User, database.QueryStatus) {
	panic("implement me")
}

func (u UserMongoDriver) Get(name string) (*user.User, database.QueryStatus) {
	panic("implement me")
}

func (u UserMongoDriver) Delete(name string) database.QueryStatus {
	panic("implement me")
}

func (u UserMongoDriver) Update(name string) database.QueryStatus {
	panic("implement me")
}
