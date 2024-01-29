package user

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/secondary_adapters/models"
)

type IUserRepository interface {
	CreateUser(string, string, string) (*int, error)
}

type UserRepository struct{}

func (u *UserRepository) CreateUser(name string, email string, password string) (*int, error) {

	var user models.User
	user.Name = name
	user.Email = email
	user.Password = password

	result := config.DB.Create(&user)

	if result.Error != nil {
		err := new(ErrorCreatingRecordInDB)
		err.Model = "user"
		return nil, err
	}

	return &user.ID, nil
}
