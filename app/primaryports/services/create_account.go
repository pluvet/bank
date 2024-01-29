package services

import (
	"github.com/pluvet/bank/app/secondaryports/eventpublisher"
	"github.com/pluvet/bank/app/secondaryports/repositories"
)

type CreateAccountService struct {
	repo           repositories.IAccountRepository
	eventPublisher eventpublisher.EventPublisher
}

func NewCreateAccountService(repo repositories.IAccountRepository, eventPublisher eventpublisher.EventPublisher) CreateAccountService {
	createAccountService := new(CreateAccountService)
	createAccountService.repo = repo
	if eventPublisher != nil {
		createAccountService.eventPublisher = eventPublisher
	}
	return *createAccountService
}

func (u *CreateAccountService) Execute(userID int) (*int, error) {

	accountID, err := u.repo.CreateAccount(userID)

	if err != nil {
		return nil, err
	}

	return accountID, nil
}
