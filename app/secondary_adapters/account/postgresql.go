package account

import (
	"github.com/pluvet/go-bank/app/secondary_adapters/user"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID      int     `json:"id" gorm:"primary_key"`
	Balance float32 `json:"balance" default:"0"`
	UserID  uint64
	User    user.User
}
