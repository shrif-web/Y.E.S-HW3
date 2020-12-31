package controller

import (
	"yes-blog/pkg/database"
)

/* 	singleton object for postController
this controller task is to perform CRUD for post.Post model
it takes a dbDriver implementing database.PostDBDriver and
speaks to the database with the dbDriver for performing
the CRUD
*/

type postController struct {
	dbDriver database.PostDBDriver
}

var pc *postController

func init() {
	pc = &postController{}
}

func GetPostController() *postController {
	return pc
}

func (p *postController) SetDBDriver(dbDriver database.PostDBDriver) {
	pc.dbDriver = dbDriver
}
