package services

import (
	"fmt"

	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/secondaryports/eventpublisher"
	"github.com/pluvet/bank/app/secondaryports/repositories"
)

type DepositAccountService struct {
	repo           repositories.IAccountRepository
	eventPublisher eventpublisher.EventPublisher
}

func NewDepositAccountService(repo repositories.IAccountRepository, eventPublisher eventpublisher.EventPublisher) DepositAccountService {
	depositAccountService := new(DepositAccountService)
	depositAccountService.repo = repo
	if eventPublisher != nil {
		depositAccountService.eventPublisher = eventPublisher
	}
	return *depositAccountService
}

func (a *DepositAccountService) Execute(accountID int, amount float32) (*float32, error) {
	account, findError := a.repo.FindAccount(accountID)
	if findError != nil {
		return nil, findError
	}

	account.Deposit(float32(amount))

	updateError := a.repo.UpdateAccount(account)
	if updateError != nil {
		return nil, updateError
	}

	var eventAccountBalanceIncreased = core.NewEventAccountBalanceIncreased(amount, account.Balance)
	eventWasPublished := a.eventPublisher.NewEvent(eventAccountBalanceIncreased)
	if !eventWasPublished {
		fmt.Printf("eventAccountBalanceIncreased was not published")
	}
	return &account.Balance, nil
}
