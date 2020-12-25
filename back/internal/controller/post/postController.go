package controller


import "yes-blog/pkg/database"

type PostController struct{
	dbDriver database.UserDBDriver
}

var postC *PostController

func init(){
	postC=&PostController{}
}

func GetPostController() *PostController {
	return postC
}

func SetDBDriver(dbDriver database.UserDBDriver){
	postC.dbDriver=dbDriver
}