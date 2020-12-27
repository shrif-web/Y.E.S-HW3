package database

import (
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database/status"
)


type UserDBDriver interface {
	Insert(user *user.User) status.QueryStatus
	Get(name *string) (*user.User,status.QueryStatus)
	Delete(name *string) status.QueryStatus
	Update(user *user.User)	status.QueryStatus
	GetAll(start,amount int64) ([]*user.User,status.QueryStatus)
}