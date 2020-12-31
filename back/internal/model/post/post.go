package post

import (
	"errors"
	"sort"
	"strconv"
	"time"
	"yes-blog/internal"
)

const ConstructorErrMsg string = "both title & content fields are empty"
const ArgErrMsg string = "POST empty exception"

var upid = 0

type Post struct {
	ID        string `json:"id" bson:"id"`
	Title     string `json:"title" bson:"title"`
	Body      string `json:"body" bson:"body"`
	Author    string `json:"author" bson:"author"`
	TimeStamp int64  `json:"timeStamp" bson:"timeStamp"`
}

func NewPost(title, body, authorID string) (*Post, error) {
	if internal.IsAllEmpty(title, body) || internal.IsEmpty(authorID) {
		return nil, errors.New(ConstructorErrMsg)
	}
	defer func() { upid++ }()
	return &Post{
		ID:        strconv.Itoa(upid),
		Title:     title,
		Body:      body,
		Author:    authorID,
		TimeStamp: time.Now().Unix(),
	}, nil
}

func NewRawPost(id, title, body, authorID string, timeStamp int64) (*Post, error) {
	if internal.IsAllEmpty(title, body) {
		return nil, errors.New(ConstructorErrMsg)
	}
	return &Post{
		ID:        id,
		Title:     title,
		Body:      body,
		Author:    authorID,
		TimeStamp: timeStamp,
	}, nil
}

func Sort(arr []*Post) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].TimeStamp >= arr[j].TimeStamp
	})
}

func Find(arr []*Post, id string) (*Post, int,  bool) {
	for i, p := range arr {
		if p.ID == id {
			return p, i, true
		}
	}
	return nil, -1, false
}

func (p *Post) setTitle(title string) error {
	if internal.IsAllEmpty(p.Body, title) {
		return errors.New(ArgErrMsg)
	}
	p.Title = title
	return nil
}

func (p *Post) SetBody(body string) error {
	if internal.IsAllEmpty(body, p.Title) {
		return errors.New(ArgErrMsg)
	}
	p.Body = body
	return nil
}

func (p *Post) UpdateTimeStamp() {
	p.TimeStamp = time.Now().Unix()
}
