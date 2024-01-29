package services

import (
	"fmt"

	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/secondaryports/eventpublisher"
	"github.com/pluvet/bank/app/secondaryports/repositories"
)

type CreateUserService struct {
	repo           repositories.IUserRepository
	eventPublisher eventpublisher.EventPublisher
}

func NewCreateUserService(repo repositories.IUserRepository, eventPublisher eventpublisher.EventPublisher) *CreateUserService {
	createUserService := new(CreateUserService)
	createUserService.repo = repo
	createUserService.eventPublisher = eventPublisher
	return createUserService
}

func (u *CreateUserService) Execute(name string, email string, password string) (*int, error) {

	userID, err := u.repo.CreateUser(name, email, password)

	if err != nil {
		return nil, err
	}

	var eventUserCreated = core.NewEventUserCreated(*userID)
	eventWasPublished := u.eventPublisher.NewEvent(eventUserCreated)

	if !eventWasPublished {
		fmt.Printf("eventUserCreated was not published")
	}

	return userID, nil
}
