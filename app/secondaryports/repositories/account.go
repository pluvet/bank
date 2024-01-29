package repositories

import (
	"github.com/pluvet/bank/app/core"
)

type IAccountRepository interface {
	CreateAccount(int) (*int, error)
	FindAccount(int) (*core.Account, error)
	UpdateAccount(account *core.Account) error
}
