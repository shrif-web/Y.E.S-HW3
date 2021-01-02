package database

import (
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database/status"
)

type UserDBDriver interface {
	Delete(name *string) status.QueryStatus
	Insert(user *user.User) status.QueryStatus
	Get(name *string) (*user.User, status.QueryStatus)
	Update(target string, user *user.User) status.QueryStatus
	GetAll(start, amount int64) ([]*user.User, status.QueryStatus)
	Replace(target *string,toBe *user.User) status.QueryStatus
}
