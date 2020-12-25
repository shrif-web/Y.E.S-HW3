package database

import "yes-blog/graph/model"


type PostDB interface {
	Insert() (*model.User,QueryStatus)
	Get(name string) (*model.User,QueryStatus)
	Delete(name string) QueryStatus
	Update(name string)	QueryStatus
}