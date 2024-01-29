package handlers

import (
	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/primaryports/services"
	"github.com/pluvet/bank/app/secondaryadapters/eventpublisher"
	"github.com/pluvet/bank/app/secondaryadapters/repositories/account"
)

type AccountHandler struct {
}

func (a *AccountHandler) HandleEvent(event core.Event) error {
	eventUserCreated, ok := event.(*core.EventUserCreated)

	if !ok {
		err := new(ErrorEventIsNotSupported)
		return err
	}

	accountService := services.NewCreateAccountService(new(account.AccountRepository), new(eventpublisher.FakeEventPublisher))
	_, err := accountService.Execute(eventUserCreated.UserID)
	if err != nil {
		return err
	}
	return nil
}
