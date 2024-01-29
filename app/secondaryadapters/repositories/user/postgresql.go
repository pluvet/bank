package user

import (
	"github.com/pluvet/bank/app/infraestructure/database"
	"github.com/pluvet/bank/app/secondaryadapters/models"
	"github.com/pluvet/bank/app/secondaryadapters/repositories/errors"
)

type UserRepository struct{}

func (u *UserRepository) CreateUser(name string, email string, password string) (*int, error) {

	var user models.User
	user.Name = name
	user.Email = email
	user.Password = password

	result := database.DB.Create(&user)

	if result.Error != nil {
		err := new(errors.ErrorCreatingRecordInDB)
		return nil, err
	}

	return &user.ID, nil
}
