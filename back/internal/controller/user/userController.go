package controller

import (
	"errors"
	"yes-blog/graph/model"
	"yes-blog/pkg/database"
	"yes-blog/pkg/database/status"
)

type UserController struct{
	dbDriver database.UserDBDriver
}

func (c *UserController) GetAll() ([]*model.User, error) {
	all, err := c.dbDriver.GetAll(-1)
	if err == status.FAILED{
		return nil,errors.New("couldn't fetch the users required")
	}
	var result[]*model.User
	for _,blogUser := range all {
		var graphUser = &model.User{
			ID:    blogUser.ID,
			Name:  blogUser.Name,
			Posts: []*model.Post{},
		}
		for _,p := range blogUser.Posts{
			graphUser.Posts = append(graphUser.Posts,&model.Post{
				ID:        p.ID,
				Auther:    graphUser,
				Timestamp: p.TimeStamp,
				Body:      p.Body,
				Title:     p.Title,
			})
		}
		result = append(result,graphUser)
	}
	return result,nil
}

var userC *UserController

func setDBDriver(dbDriver database.UserDBDriver){
	userC.dbDriver=dbDriver
}

func init(){
	userC=&UserController{}
}

func GetUserController() *UserController {
	return userC

}