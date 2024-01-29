package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/primaryadapters/publisher"
	"github.com/pluvet/bank/app/primaryports/services"
	"github.com/pluvet/bank/app/secondaryadapters/repositories/account"
	"github.com/pluvet/bank/app/secondaryadapters/repositories/errors"
)

type AccountBalanceChangeInputDTO struct {
	Amount float32
}

type AccountBalanceChangeOutputDTO struct {
	Balance float32
}

func AccountDeposit(c *gin.Context) {
	var accountBalanceChangeDTO AccountBalanceChangeInputDTO
	c.BindJSON(&accountBalanceChangeDTO)
	stringID := c.Param("id")
	accountID, idError := strconv.Atoi(stringID)
	if idError != nil {
		c.JSON(400, "id is not an int value")
	}
	accountDepositService := services.NewDepositAccountService(new(account.AccountRepository), publisher.GetEventPublisher())
	newBalance, err := accountDepositService.Execute(accountID, accountBalanceChangeDTO.Amount)
	if err != nil {
		switch err.(type) {
		case *errors.ErrorFindingOneRecordInDB:
			c.JSON(404, err.Error())
		case *errors.ErrorUpdatingRecordInDB:
			c.JSON(500, err.Error())
		default:
			c.JSON(500, "Internal Server Error")
		}
	}
	var output AccountBalanceChangeOutputDTO
	output.Balance = *newBalance
	c.JSON(200, output)
}

func AccountWithdraw(c *gin.Context) {
	var accountBalanceChangeDTO AccountBalanceChangeInputDTO
	c.BindJSON(&accountBalanceChangeDTO)
	stringID := c.Param("id")
	accountID, idError := strconv.Atoi(stringID)
	if idError != nil {
		c.JSON(400, "id is not an int value")
	}
	accountWithdrawService := services.NewWithdrawAccountService(new(account.AccountRepository), publisher.GetEventPublisher())
	newBalance, err := accountWithdrawService.Execute(accountID, accountBalanceChangeDTO.Amount)
	if err != nil {
		switch err.(type) {
		case *errors.ErrorFindingOneRecordInDB:
			c.JSON(404, err.Error())
		case *errors.ErrorUpdatingRecordInDB:
			c.JSON(500, err.Error())
		case *core.WithdrawError:
			c.JSON(400, err.Error())
		default:
			c.JSON(500, "Internal Server Error")
		}
	}
	var output AccountBalanceChangeOutputDTO
	output.Balance = *newBalance
	c.JSON(200, output)
}
