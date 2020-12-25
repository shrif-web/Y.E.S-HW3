package controller


import "yes-blog/pkg/database"

type PostController struct{
	dbDriver database.UserDBDriver
}

func NewPostController(dbDriver database.UserDBDriver) *PostController {
	return &PostController{
		dbDriver: dbDriver,
	}
}