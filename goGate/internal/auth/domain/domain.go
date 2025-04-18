package domain

type User struct {
	ID       int64
	Username string
	Password string
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
	Create(user *User) error
}
