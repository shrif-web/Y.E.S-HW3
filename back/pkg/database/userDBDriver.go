package database

import "yes-blog/internal/model/user"


type UserDBDriver interface {
	Insert() (*user.User,QueryStatus)
	Get(name string) (*user.User,QueryStatus)
	Delete(name string) QueryStatus
	Update(name string)	QueryStatus
}