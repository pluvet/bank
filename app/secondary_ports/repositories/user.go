package repositories

type UserRepository interface {
	CreateUser(string, string, string) (*int, error)
}
