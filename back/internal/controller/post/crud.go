package controller

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"yes-blog/internal/model/post"
	"yes-blog/pkg/database/status"
)

/*	the actual implementation of CRUD for postController is here
	we also added getAll method due to our certain needs
*/
/*	we use status.QueryStatus as a statusCode for our controllers
	we use status.FAILED to return a failed status and
	status.SUCCESSFUL to return a successful status (obviously)
*/

// create model.Post entry then add into DB
func (p *postController) WritePost(title, body, authorID string) (*post.Post, error) {
	newPost, err := post.NewPost(title, body, authorID)
	if err != nil {
		return nil, err
	}

	stat := p.dbDriver.Insert(newPost)
	if stat == status.FAILED {
		return nil, errors.New("dataBase Failed to Insert A Post")
	}
	return newPost, nil
}

// edit the post in DB
func (p *postController) EditPost(postID, title, body, authorID string) (string, error) {
	upPost, err := post.NewRawPost(postID, title, body, authorID, time.Now().Unix())
	if err != nil {
		return fmt.Sprint(err), err
	}
	stat, err := p.dbDriver.Update(upPost)
	if stat == status.FAILED {
		return "the post couldn't edit", err
	}
	return "the edited successfully", nil
}

// delete the post from DB
func (p *postController) DeletePost(postID string) (string, error) {
	stat := p.dbDriver.Delete(postID)
	if stat == status.FAILED {
		return "the post couldn't delete", errors.New("dataBase Failed to Delete The Post " + postID)
	}
	return "the post deleted successfully", nil
}

// get the post specified with id from DB
func (p *postController) GetPost(postID string) (*post.Post, error) {
	pst, stat := p.dbDriver.Get(postID)
	if stat == status.FAILED {
		return nil, errors.New("dataBase Failed to get The Post " + postID)
	}
	return pst, nil
}

// get all posts from DB with start id and size of amount
func (p *postController) GetAllPosts(startID string, amount uint64) ([]*post.Post, error) {
	sID, _ := strconv.ParseUint(startID, 10, 64)
	posts, stat := p.dbDriver.GetAll(sID, amount)
	if stat == status.FAILED {
		return nil, errors.New("dataBase Failed to get Posts from " + startID + " to " + strconv.FormatUint(amount, 10))
	}
	return posts, nil
}

// get all posts which is written by user from DB
func (p *postController) GetPostByUser(userID string) ([]*post.Post, error) {
	posts, stat := p.dbDriver.GetByUser(userID)
	if stat == status.FAILED {
		return nil, errors.New("dataBase Failed to get Posts from user " + userID)
	}
	return posts, nil
}
