package post

import (
	"time"
	"yes-blog/graph/model"
)

type Post struct {
	Title     string
	Body      string
	Author    *model.User
	TimeStamp time.Time
}

func NewPost(title, body string, author *model.User) *Post {
	return &Post{
		Title:     title,
		Body:      body,
		Author:    author,
		TimeStamp: time.Now(),
	}
}

func (p *Post) EditTitle(title string) {
	p.Title = title
}
func (p *Post) EditBody(body string) {
	p.Body = body
}
