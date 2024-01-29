package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/publisher"
	"github.com/pluvet/go-bank/app/repositories"
	"github.com/pluvet/go-bank/app/services"
)

type CreateUserInputDTO struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutputDTO struct {
	ID int
}

func CreateUser(c *gin.Context) {
	var user CreateUserInputDTO
	c.BindJSON(&user)
	userService := services.NewUserService(new(repositories.UserRepository), *publisher.GetEventPublisher())
	userID, err := userService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		switch err.(type) {
		case *repositories.ErrorCreatingRecordInDB:
			c.JSON(500, err.Error())
		default:
			c.JSON(500, "Internal Server Error")
		}
	}
	userOutputDTO := new(CreateUserOutputDTO)
	userOutputDTO.ID = *userID

	c.JSON(201, userOutputDTO)
}
