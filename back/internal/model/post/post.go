package post

import (
	"crypto/sha1"
	"errors"
	"sort"
	"strconv"
	"time"
)

const ConstructorErrMsg string = "both title & content fields are empty"
const ArgErrMsg string = "POST empty exception"

type Post struct {
	ID        string `json:"id" bson:"id"`
	Title     string `json:"title" bson:"title"`
	Body      string `json:"body" bson:"body"`
	AuthorID  string `json:"author" bson:"author"`
	TimeStamp int64  `json:"timeStamp" bson:"timeStamp"`
}

func NewPost(title, body, authorID string) (*Post, error) {
	if isAllEmpty(title, body) || isEmpty(authorID) {
		return nil, errors.New(ConstructorErrMsg)
	}
	pst := &Post{
		ID:        "",
		Title:     title,
		Body:      body,
		AuthorID:  authorID,
		TimeStamp: time.Now().Unix(),
	}
	pst.ID = pst.GetHash()
	return pst, nil
}

func NewRawPost(id, title, body, authorID string, timeStamp int64) (*Post, error) {
	if isAllEmpty(title, body) {
		return nil, errors.New(ConstructorErrMsg)
	}
	return &Post{
		ID:        id,
		Title:     title,
		Body:      body,
		AuthorID:  authorID,
		TimeStamp: timeStamp,
	}, nil
}

func Sort(arr []*Post) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].TimeStamp >= arr[j].TimeStamp
	})
}

func Find(arr []*Post, id string) (*Post, int, bool) {
	for i, p := range arr {
		if p.ID == id {
			return p, i, true
		}
	}
	return nil, -1, false
}

func (p *Post) GetHash() string {
	return string(sha1.New().Sum([]byte(p.AuthorID + "|" + strconv.FormatInt(p.TimeStamp, 10))))
}

func (p *Post) setTitle(title string) error {
	if isAllEmpty(p.Body, title) {
		return errors.New(ArgErrMsg)
	}
	p.Title = title
	return nil
}

func (p *Post) SetBody(body string) error {
	if isAllEmpty(body, p.Title) {
		return errors.New(ArgErrMsg)
	}
	p.Body = body
	return nil
}

func (p *Post) UpdateTimeStamp() {
	p.TimeStamp = time.Now().Unix()
}
