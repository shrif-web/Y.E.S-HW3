package user

import (
	"strconv"
	"yes-blog/graph/model"
	"yes-blog/internal/model/post"
)
var uuid=0
type User struct {
	id       string
	Name     string
	Password string
	Admin    bool
	Posts    []*post.Post
}

func NewUser(name, password string) (*User,error) {

	return newUser(name, password, false)
}

func NewAdmin(name, password string) (*User, error) {
	return newUser(name, password, true)
}

func newUser(name, password string, admin bool) (*User,error) {
	defer func(){uuid++}()
	// hashing password
	hashedPass,err:=hashAndSalt([]byte(password))
	if err!=nil{
		message:="internal server error: couldn't hash password"
		return nil,model.InternalServerException{Message: &message}
	}

	return &User{
		ID:       strconv.Itoa(uuid),
		Name:     name,
		Password: hashedPass,
		Admin:    admin,
		Posts: []*post.Post{},
	},nil
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

func (u *User) UpdatePassword(password string) error {
	// hashing password
	hashedPass,err:=hashAndSalt([]byte(password))

	if err!=nil{
		message:="internal server error: couldn't hash password"
		return model.InternalServerException{Message: &message}
	}

	u.Password=hashedPass
	return nil
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) Upgrade() {
	u.Admin = true
}

func (u *User) degrade() {
	u.Admin = false
}
func (u *User)getID() string {
	return u.id
}
func (u *User)setID(newId string){
	u.id=newId
}