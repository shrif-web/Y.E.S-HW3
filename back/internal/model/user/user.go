package user

import (
	"strconv"
	"yes-blog/internal/model/post"
)
var uuid=0
type User struct {
	ID		 string
	Name     string
	password string
	admin    bool
	Posts    []*post.Post
}

func NewUser(name, password string) *User {
	return newUser(name, password, false)
}

func NewAdmin(name, password string) *User {
	return newUser(name, password, true)
}

func newUser(name, password string, admin bool) *User {
	defer func(){uuid++}()
	return &User{
		ID:	strconv.Itoa(uuid),
		Name:     name,
		password: password,
		admin:    admin,
	}
}

func (u *User) AddPost(p *post.Post) *User {
	u.Posts = append(u.Posts, p)
	return u
}

func (u *User) DeletePost(id string) *User{
	for i,p := range u.Posts{
		if p.ID==id{
			u.Posts = append(u.Posts[:i], u.Posts[i+1:]...)
			return u
		}
	}
	return u
}

func (u *User) UpdatePassword(password string) *User {
	u.password = password
	return u
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) Upgrade() {
	u.admin = true
}

func (u *User) degrade() {
	u.admin = false
}
