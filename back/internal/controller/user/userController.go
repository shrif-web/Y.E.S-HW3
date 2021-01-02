package controller

import (
	"yes-blog/graph/model"
	"yes-blog/pkg/database"
	"yes-blog/pkg/jwt"
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

func (c *userController) Login(username,password string)(string,error){

	//retrieve user from data base
	blogUser, err := c.Get(&username)
	if err != nil {
		switch err.(type) {
		case *model.PostEmptyException:
			return "",err.(*model.InternalServerException)
		}
		return "",err
	}

	// check if the username and password matches
	if !blogUser.Verify(password) {
		return "",model.UserPassMissMatchException{}
	}

	// generate new token
	token, err2 := jwt.GenerateToken(blogUser.Name)
	if err2 != nil {
		return "",model.InternalServerException{}
	}
	return token, nil
}