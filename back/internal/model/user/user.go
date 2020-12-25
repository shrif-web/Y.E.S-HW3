package user

type User struct {
	Name     string
	password string
	admin    bool
}

func NewUser(name,password string) *User{
	return newUser(name, password,false)
}

func NewAdmin(name,password string) *User{
	return newUser(name, password,true)
}

func newUser(name,password string , admin bool) *User{
	return &User{
		Name: name,
		password: password,
		admin: admin,
	}
}

func (u *User)UpdatePassword(password string) {
	u.password=password
}

func (u *User)UpdateName(name string) {
	u.Name=name
}

func (u *User)Upgrade() {
	u.admin=true
}

func (u *User)degrade() {
	u.admin=false
}