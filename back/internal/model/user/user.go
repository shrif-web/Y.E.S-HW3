package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-blog/graph/model"
	"yes-blog/internal/model/post"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string
	Password string
	Email    string
	Admin    bool
	Posts    []*post.Post
}

func NewUser(name, password, email string) (*User, error) {

	return newUser(name, password, email, false)
}

func NewAdmin(name, password, email string) (*User, error) {
	return newUser(name, password, email, true)
}

func newUser(name, password, email string, admin bool) (*User, error) {
	// hashing password
	hashedPass, err := hashAndSalt([]byte(password))
	if err != nil {
		return nil, model.InternalServerException{Message: "internal server error: couldn't hash password"}
	}

	return &User{
		Name:     name,
		Password: hashedPass,
		Email:    email,
		Admin:    admin,
		Posts:    []*post.Post{},
	}, nil
}

func (u *User) AddPost(p *post.Post) *User {
	u.Posts = append(u.Posts, p)
	return u
}

func (u *User) DeletePost(id string) *User {
	for i, p := range u.Posts {
		if p.ID.Hex() == id {
			u.Posts = append(u.Posts[:i], u.Posts[i+1:]...)
			return u
		}
	}
	return u
}

func (u *User) UpdatePassword(password string) error {
	// hashing password
	hashedPass, err := hashAndSalt([]byte(password))

	if err != nil {
		return model.InternalServerException{Message: "internal server error: couldn't hash password"}
	}

	u.Password = hashedPass
	return nil
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) promote() {
	u.Admin = true
}

func (u *User) demote() {
	u.Admin = false
}

func (u *User) Verify(password string) bool {
	return CheckPasswordHash(password, u.Password)
}
