package repositories

type IUserRepository interface {
	CreateUser(string, string, string) (*int, error)
}
