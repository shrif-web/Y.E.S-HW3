package controller

import "yes-blog/pkg/database"

type UserController struct{
	dbDriver database.UserDBDriver
}

func NewUserController(dbDriver database.UserDBDriver) *UserController {
	return &UserController{
		dbDriver: dbDriver,
	}
}