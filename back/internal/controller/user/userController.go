package controller

import (
	"yes-blog/pkg/database"
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
