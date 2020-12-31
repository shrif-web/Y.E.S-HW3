package post

import (
	"errors"
	"strconv"
	"time"
	"yes-blog/internal"
)

const ConstructorErrMsg string = "POST constructor argument input exception"
const ArgErrMsg string = "POST empty exception"

var uuid = 0

type Post struct {
	id        string
	title     string
	body      string
	authorID  string
	timeStamp int64
}

func NewPost(title, body, authorID string) (*Post, error) {
	if internal.IsAllEmpty(title, body) || internal.IsEmpty(authorID) {
		return nil, errors.New(ConstructorErrMsg)
	}
	defer func() { uuid++ }()
	return &Post{
		id:        strconv.Itoa(uuid),
		title:     title,
		body:      body,
		authorID:  authorID,
		timeStamp: time.Now().Unix(),
	}, nil
}

func NewRawPost(id, title, body, authorID string, timeStamp int64) (*Post, error) {
	if internal.IsAllEmpty(title, body) {
		return nil, errors.New(ConstructorErrMsg)
	}
	return &Post{
		id:        id,
		title:     title,
		body:      body,
		authorID:  authorID,
		timeStamp: timeStamp,
	}, nil
}

func (p *Post) GetID() string {
	return p.id
}

func (p *Post) GetTitle() string {
	return p.title
}
func (p *Post) GetBody() string {
	return p.body
}
func (p *Post) GetAuthorID() string {
	return p.authorID
}
func (p *Post) GetTimeStamp() int64 {
	return p.timeStamp
}

func (p *Post) setTitle(title string) error {
	if internal.IsAllEmpty(p.body, title) {
		return errors.New(ArgErrMsg)
	}
	p.title = title
	return nil
}

func (p *Post) SetBody(body string) error {
	if internal.IsAllEmpty(body, p.title) {
		return errors.New(ArgErrMsg)
	}
	p.body = body
	return nil
}

func (p *Post) UpdateTimeStamp() {
	p.timeStamp = time.Now().Unix()
}
