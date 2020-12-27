package post

import (
	"strconv"
	"time"
)

var uuid = 0

type Post struct {
	ID        string
	Title     string
	Body      string
	AuthorID  string
	TimeStamp int64
}

func NewPost(title, body string, authorID string) *Post {
	defer func() { uuid++ }()
	return &Post{
		ID:        strconv.Itoa(uuid),
		Title:     title,
		Body:      body,
		AuthorID:  authorID,
		TimeStamp: time.Now().Unix(),
	}
}

func (p *Post) EditTitle(title string) {
	p.Title = title
}
func (p *Post) EditBody(body string) {
	p.Body = body
}
