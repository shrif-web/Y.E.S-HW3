package controller

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-blog/graph/model"
	userController "yes-blog/internal/controller/user"
	"yes-blog/internal/model/post"
	"yes-blog/internal/model/user"
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
func (p *postController) CreatePost(title, body, authorName string) (*post.Post, *user.User, error) {

	usr, err := userController.GetUserController().Get(&authorName)
	if err != nil {
		return nil, nil, &model.NoUserFoundException{Message: err.Error()}
	}

	newPost, errr := post.NewPost(title, body, usr.ID.Hex())
	if errr != nil {
		message := errr.Error()
		return nil, nil, &model.PostEmptyException{Message: message}
	}

	err = p.dbDriver.Insert(newPost, authorName)
	err = CastDBEToGQLE(err)
	if err != nil {
		return nil, nil, err
	}
	return newPost, usr, nil
}

// edit the post in DB
func (p *postController) UpdatePost(postID, title, body, operator, authorName string) (string, error) {
	c, err := userController.GetUserController().CanOperate(operator, authorName)
	if err != nil {
		switch err.(type) {
		case *model.NoUserFoundException:
			return err.Error(), &model.NoUserFoundException{Message: err.Error()}
		default:
			return err.Error(), &model.InternalServerException{Message: err.Error()}
		}
	}
	if !c {
		return "you are not allowed to update post@" + postID, &model.UserNotAllowedException{Message: "you are not allowed to update post@" + postID}
	}

	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err.Error(), &model.InternalServerException{Message: err.Error()}
	}
	upPost, err := post.NewRawPost(id, title, body, authorName, time.Now().Unix())
	if err != nil {
		return err.Error(), &model.InternalServerException{Message: err.Error()}
	}
	err = p.dbDriver.Update(upPost)
	err = CastDBEToGQLE(err)
	if err != nil {
		return "the post couldn't edit", err
	}
	return "the post edited successfully", nil
}

// delete the post from DB
func (p *postController) DeletePost(postID string, operator, authorName string) (string, error) {
	c, err := userController.GetUserController().CanOperate(operator, authorName)
	if err != nil {
		switch err.(type) {
		case *model.NoUserFoundException:
			return err.Error(), &model.NoUserFoundException{Message: err.Error()}
		default:
			return err.Error(), &model.InternalServerException{Message: err.Error()}
		}
	}
	if !c {
		return "you are not allowed to delete post@" + postID, &model.UserNotAllowedException{Message: "you are not allowed to delete post@" + postID}
	}

	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return fmt.Sprint(err), err
	}
	err = p.dbDriver.Delete(id, authorName)
	err = CastDBEToGQLE(err)
	if err != nil {
		return "the post couldn't delete", err
	}
	return "the post deleted successfully", nil
}

// get the post specified with id from DB
func (p *postController) GetPost(postID string) (*post.Post, *user.User, error) {
	id, errr := primitive.ObjectIDFromHex(postID)
	if errr != nil {
		return nil, nil, errr
	}
	pst, usr, err := p.dbDriver.Get(id)
	if err != nil {
		return nil, nil, err
	}
	return pst, usr, nil
}

// get all posts from DB with start id and size of amount
func (p *postController) GetAllPosts(startIndex, amount int) ([]*post.Post, []*user.User, error) {
	psts, usrs, err := p.dbDriver.GetAll(startIndex, amount)
	if err != nil {
		return nil, nil, err
	}
	return psts, usrs, nil
}

// get all posts which is written by user from DB
func (p *postController) GetPostByUser(userName string) ([]*post.Post, *user.User, error) {
	psts, usr, err := p.dbDriver.GetByUser(userName)
	if err != nil {
		return nil, nil, err
	}
	return psts, usr, nil
}

// casting database errors to model.graphQL exceptions
func CastDBEToGQLE(err error) error {
	if err != nil {
		switch err.(type) {
		case *DBException.UserNotFound:
			return &model.NoUserFoundException{Message: err.Error()}
		case *DBException.PostNotFound:
			return &model.PostNotFoundException{Message: err.Error()}
		case *DBException.UserNotAllowed:
			return &model.UserNotAllowedException{Message: err.Error()}
		case *DBException.InternalDBError:
			return &model.InternalServerException{Message: err.Error()}
		}
	}
	return nil
}
