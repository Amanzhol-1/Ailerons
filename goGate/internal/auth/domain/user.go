package domain

type User struct {
	ID       int64
	Username string
	Password string
	Role     string // "customer", "driver"
}
