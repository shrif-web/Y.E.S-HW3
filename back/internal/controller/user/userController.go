package controller

import (
	"errors"
	"yes-blog/graph/model"
	"yes-blog/pkg/database"
	"yes-blog/pkg/database/status"
)

type UserController struct {
	dbDriver database.UserDBDriver
}

var userC *UserController

func (c *UserController) SetDBDriver(dbDriver database.UserDBDriver) {
	userC.dbDriver = dbDriver
}




func init() {
	userC = &UserController{}
}

func GetUserController() *UserController {
	return userC
}
