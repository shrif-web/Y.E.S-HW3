package controller

import (
	"errors"
	"yes-blog/graph/model"
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database/status"
)

/*	the actual implementation of CRUD for userController is here
	we also added getAll method due to our certain needs
*/
/*	we use status.QueryStatus as a statusCode for our controllers
	we use status.FAILED to return a failed status and
	status.SUCCESSFUL to return a successful status (obviously)
*/
func (c *userController) GetAll(start, amount int64) ([]*user.User, error) {
	all, err := c.dbDriver.GetAll(start, amount)
	if err == status.FAILED {
		return nil, errors.New("couldn't fetch the users required")
	}
	return all, nil
}

func (c *userController) Get(name *string) (*user.User, error) {
	if blogUser, stat := c.dbDriver.Get(name); stat == status.FAILED {
		message:="couldn't find the requested user"
		return nil, model.NoUserFoundException{Message: &message}
	} else {
		return blogUser, nil
	}
}

func (c *userController) Delete(name *string) error {
	if stat := c.dbDriver.Delete(name); stat == status.FAILED {
		message :="couldn't delete the user"
		return model.InternalServerException{Message: &message}
	} else {
		return nil
	}
}

func (c *userController) Update(target string, toBe model.ToBeUser) (*user.User,error) {

	// filling the update fields of the user
	var blogUser = user.User{}
	if toBe.Username != nil {
		blogUser.Name = *(toBe.Username)
	}
	if toBe.Password != nil {
		blogUser.UpdatePassword(*(toBe.Password))
	}

	// updating the database
	if stat := c.dbDriver.Update(target, &blogUser); stat == status.FAILED {

		// checking if the target user exists
		_, stat2 := c.dbDriver.Get(&target)
		if stat2 == status.FAILED {
			message:="target Doesnt exist"
			return nil, model.NoUserFoundException{Message: &message}
		}
		// no clue why query failed
		message:="couldn't update the user"
		return nil,model.InternalServerException{Message: &message}
	} else {
		return &blogUser,nil
	}
}

func (c *userController) Create(name, password string) (*user.User, error) {

	// checking for duplicate username
	if _,stat := c.dbDriver.Get(&name); stat == status.SUCCESSFUL {
		return nil, model.DuplicateUsernameException{}
	}

	// creating new user Object to insert in to the data base
	newUser,err := user.NewUser(name, password)
	if err!=nil{
		return nil, err
	}
	// inserting into the database
	if stat := c.dbDriver.Insert(newUser); stat == status.FAILED {
		message:="couldn't create user"
		return &user.User{},model.InternalServerException{Message: &message}
	} else {
		return newUser, nil
	}
}
