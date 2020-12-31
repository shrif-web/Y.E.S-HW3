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
		return nil, errors.New("couldn't find the requested user")
	} else {
		return blogUser, nil
	}
}

func (c *userController) Delete(name *string) error {
	if stat := c.dbDriver.Delete(name); stat == status.FAILED {
		return errors.New("couldn't delete the user")
	} else {
		return nil
	}
}

func (c *userController) Update(target string, toBe model.ToBeUser) error {
	//todo return error if target doesnt exist
	var blogUser = user.User{}
	if toBe.Username != nil {
		blogUser.Name = *(toBe.Username)
	}
	if toBe.Password != nil {
		blogUser.UpdatePassword(*(toBe.Password))
	}
	if stat := c.dbDriver.Update(target, &blogUser); stat == status.FAILED {
		return errors.New("couldn't update the user")
	} else {
		return nil
	}
}

func (c *userController) Create(name, password string) (*user.User, error) {
	hashedPass,err:=hashAndSalt([]byte(password))
	if err!=nil{
		return &user.User{},errors.New("internal server error: couldn't hash password")
	}
	newUser := user.NewUser(name, hashedPass)
	if _,stat := c.dbDriver.Get(&name); stat == status.SUCCESSFUL {
		return nil, errors.New("user with this name already exist")
	}
	if stat := c.dbDriver.Insert(newUser); stat == status.FAILED {
		return nil, errors.New("couldn't create user")
	} else {
		return newUser, nil
	}
}
