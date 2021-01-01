package controller

import (
	"fmt"
	"time"
	"yes-blog/graph/model"
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database/DBException"
)

/*	the actual implementation of CRUD for postController is here
	we also added getAll method due to our certain needs
*/
/*	we use status.QueryStatus as a statusCode for our controllers
	we use status.FAILED to return a failed status and
	status.SUCCESSFUL to return a successful status (obviously)
*/

// create model.Post entry then add into DB
func (p *postController) CreatePost(title, body, authorName string) (*post.Post, error) {
	newPost, err := post.NewPost(title, body, authorName)
	if err != nil {
		message := err.Error()
		return nil, &model.PostEmptyException{Message: &message}
	}
	err = p.dbDriver.Insert(newPost)
	err = CastDBEToGQLE(err)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

// edit the post in DB
func (p *postController) UpdatePost(postID, title, body, authorName string) (string, error) {
	upPost, err := post.NewRawPost(postID, title, body, authorName, time.Now().Unix())
	if err != nil {
		return fmt.Sprint(err), err
	}
	err = p.dbDriver.Update(upPost)
	err = CastDBEToGQLE(err)
	if err != nil {
		return "the post couldn't edit", err
	}
	return "the edited successfully", nil
}

// delete the post from DB
func (p *postController) DeletePost(postID string, authorName string) (string, error) {
	err := p.dbDriver.Delete(postID, authorName)
	err = CastDBEToGQLE(err)
	if err != nil {
		return "the post couldn't delete", err
	}
	return "the post deleted successfully", nil
}

// get the post specified with id from DB
func (p *postController) GetPost(postID string) (*post.Post, error) {
	pst, err := p.dbDriver.Get(postID)
	if err != nil {
		return nil, err
	}
	return pst, nil
}

// get all posts from DB with start id and size of amount
func (p *postController) GetAllPosts(startID string, amount int) ([]*post.Post, error) {
	psts, err := p.dbDriver.GetAll(startID, amount)
	if err != nil {
		return nil, err
	}
	return psts, nil
}

// get all posts which is written by user from DB
func (p *postController) GetPostByUser(userName string) ([]*post.Post, error) {
	psts, err := p.dbDriver.GetByUser(userName)
	if err != nil {
		return nil, err
	}
	return psts, nil
}

// casting database errors to model.graphQL exceptions
func CastDBEToGQLE(err error) error {
	if err != nil {
		message := err.Error()
		switch err.(type) {
		case *DBException.UserNotFound:
			return &model.NoUserFoundException{Message: &message}
		case *DBException.PostNotFound:
			return &model.PostNotFoundException{Message: &message}
		case *DBException.UserNotAllowed:
			return &model.UserNotAllowedException{Message: &message}
		case *DBException.InternalDBError:
			return &model.InternalServerException{Message: &message}
		}
	}
	return nil
}
