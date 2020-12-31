package controller

import (
	"yes-blog/pkg/database"
)

/* 	singleton object for userController
	this controller task is to perform CRUD for user.User model
	it takes a dbDriver implementing database.UserDBDriver and
	speaks to the database with the dbDriver for performing
	the CRUD
*/
type userController struct {
	dbDriver database.UserDBDriver
}

var userC *userController

func (c *userController) SetDBDriver(dbDriver database.UserDBDriver) {
	userC.dbDriver = dbDriver
}

func init() {
	userC = &userController{}
}

func GetUserController() *userController {
	return userC
}
