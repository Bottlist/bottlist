package types

type UserType int

const (
	Client = iota
	Shop
)

type User struct {
	ID       int
	UserType UserType
}
