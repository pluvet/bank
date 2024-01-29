package account

import (
	"errors"

	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/infraestructure/database"
	"github.com/pluvet/bank/app/secondaryadapters/models"
	repoErrors "github.com/pluvet/bank/app/secondaryadapters/repositories/errors"
	"gorm.io/gorm"
)

type AccountRepository struct{}

func (a *AccountRepository) CreateAccount(userID int) (*int, error) {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(userID)
	result := database.DB.Create(&account)

	if result.Error != nil {
		err := new(repoErrors.ErrorCreatingRecordInDB)
		return nil, err
	}

	return &account.ID, nil
}

func (a *AccountRepository) FindAccount(accountID int) (*core.Account, error) {
	var account *core.Account
	dbErr := database.DB.Where("id = ?", accountID).First(&account).Error
	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		err := new(repoErrors.ErrorFindingOneRecordInDB)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	return account, nil
}

func (a *AccountRepository) UpdateAccount(account *core.Account) error {
	result := database.DB.Save(account)
	if result.Error != nil {
		err := new(repoErrors.ErrorUpdatingRecordInDB)
		return err
	}
	return nil
}
