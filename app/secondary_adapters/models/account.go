package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID      int     `json:"id" gorm:"primary_key"`
	Balance float32 `json:"balance" default:"0"`
	UserID  uint64
	User    User
}
