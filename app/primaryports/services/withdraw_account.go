package services

import (
	"fmt"

	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/secondaryports/eventpublisher"
	"github.com/pluvet/bank/app/secondaryports/repositories"
)

type WithdrawAccountService struct {
	repo           repositories.IAccountRepository
	eventPublisher eventpublisher.EventPublisher
}

func NewWithdrawAccountService(repo repositories.IAccountRepository, eventPublisher eventpublisher.EventPublisher) WithdrawAccountService {
	withdrawAccountService := new(WithdrawAccountService)
	withdrawAccountService.repo = repo
	if eventPublisher != nil {
		withdrawAccountService.eventPublisher = eventPublisher
	}
	return *withdrawAccountService
}

func (a *WithdrawAccountService) Execute(accountID int, amount float32) (*float32, error) {
	account, findError := a.repo.FindAccount(accountID)
	if findError != nil {
		return nil, findError
	}

	withdrawError := account.Withdraw(amount)
	if withdrawError != nil {
		return nil, withdrawError
	}

	updateError := a.repo.UpdateAccount(account)
	if updateError != nil {
		return nil, updateError
	}

	var eventAccountBalanceDecreased = core.NewEventAccountBalanceDecreased(amount, account.Balance)
	eventWasPublished := a.eventPublisher.NewEvent(eventAccountBalanceDecreased)
	if !eventWasPublished {
		fmt.Printf("eventAccountBalanceDecreased was not published")
	}
	return &account.Balance, nil
}
