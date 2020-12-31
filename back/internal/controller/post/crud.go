package controller

import (
	"fmt"
	"time"
	"yes-blog/internal/model/post"
)

/*	the actual implementation of CRUD for postController is here
	we also added getAll method due to our certain needs
*/
/*	we use status.QueryStatus as a statusCode for our controllers
	we use status.FAILED to return a failed status and
	status.SUCCESSFUL to return a successful status (obviously)
*/

// create model.Post entry then add into DB
func (p *postController) WritePost(title, body, authorName string) (*post.Post, error) {
	newPost, err := post.NewPost(title, body, authorName)
	if err != nil {
		return nil, err
	}

	return newPost, p.dbDriver.Insert(newPost)
}

// edit the post in DB
func (p *postController) EditPost(postID, title, body, authorName string) (string, error) {
	upPost, err := post.NewRawPost(postID, title, body, authorName, time.Now().Unix())
	if err != nil {
		return fmt.Sprint(err), err
	}
	err = p.dbDriver.Update(upPost)
	if err != nil {
		return "the post couldn't edit", err
	}
	return "the edited successfully", nil
}

// delete the post from DB
func (p *postController) DeletePost(postID string, authorName string) (string, error) {
	err := p.dbDriver.Delete(postID, authorName)
	if err != nil {
		return "the post couldn't delete", err
	}
	return "the post deleted successfully", nil
}

// get the post specified with id from DB
func (p *postController) GetPost(postID string) (*post.Post, error) {
	return p.dbDriver.Get(postID)
}

// get all posts from DB with start id and size of amount
func (p *postController) GetAllPosts(startID string, amount int) ([]*post.Post, error) {
	return p.dbDriver.GetAll(startID, amount)
}

// get all posts which is written by user from DB
func (p *postController) GetPostByUser(userName string) ([]*post.Post, error) {
	return p.dbDriver.GetByUser(userName)
}
