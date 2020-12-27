package controller

import (
	"errors"
	"yes-blog/graph/model"
	"yes-blog/pkg/database/status"
)

func (c *UserController) GetAll(start,amount int64) ([]*model.User, error) {
	all, err := c.dbDriver.GetAll(start,amount)
	if err == status.FAILED {
		return nil, errors.New("couldn't fetch the users required")
	}

	return reformatUsers(all), nil
}

func (c *UserController) Get(name *string) (*model.User, error) {
	if blogUser, stat := c.dbDriver.Get(name); stat==status.FAILED{
		return nil,errors.New("couldn't find the requested user")
	}else{
		return reformatUser(blogUser),nil
	}
}

func (c *UserController) Delete(name *string) error{
	if stat:= c.dbDriver.Delete(name); stat==status.FAILED{
		return errors.New("couldn't delete the user")
	}else{
		return nil
	}
}
