package infraestructure

import (
	"github.com/pluvet/go-bank/app/secondary_adapters/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	url := "postgres://postgres:123456@localhost:5434/bank"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Account{})
	DB = db
	DB.Statement.RaiseErrorOnNotFound = true
}
