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

func (u *User) AddPost(p *post.Post) {
	u.Posts = append(u.Posts, p)
}

func (u *User) DeletePost(id string){
	for i,p := range u.Posts{
		if p.ID==id{
			u.Posts = append(u.Posts[:i], u.Posts[i+1:]...)
			return
		}
	}
}

func (u *User) UpdatePassword(password string) {
	u.password = password
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
