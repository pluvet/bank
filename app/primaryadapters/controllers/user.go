package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/bank/app/primaryadapters/publisher"
	"github.com/pluvet/bank/app/primaryports/services"
	"github.com/pluvet/bank/app/secondaryadapters/repositories/errors"
	"github.com/pluvet/bank/app/secondaryadapters/repositories/user"
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
	var userInputDTO CreateUserInputDTO
	c.BindJSON(&userInputDTO)
	userService := services.NewCreateUserService(new(user.UserRepository), publisher.GetEventPublisher())
	userID, err := userService.Execute(userInputDTO.Name, userInputDTO.Email, userInputDTO.Password)
	if err != nil {
		switch err.(type) {
		case *errors.ErrorCreatingRecordInDB:
			c.JSON(500, err.Error())
		default:
			c.JSON(500, "Internal Server Error")
		}
	}
	userOutputDTO := new(CreateUserOutputDTO)
	userOutputDTO.ID = *userID

	c.JSON(201, userOutputDTO)
}
