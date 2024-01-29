package publisher

import (
	"github.com/pluvet/bank/app/core"
	"github.com/pluvet/bank/app/primaryadapters/handlers"
	"github.com/pluvet/bank/app/secondaryadapters/eventpublisher"
)

var eventPublisher *eventpublisher.EventPublisher

func GetEventPublisher() *eventpublisher.EventPublisher {
	if eventPublisher == nil {
		var accountHandlers = map[string][]eventpublisher.Handler{
			new(core.EventUserCreated).GetName(): {new(handlers.AccountHandler)},
		}
		eventPublisher = eventpublisher.NewEventPublisher(accountHandlers)
	}
	return eventPublisher
}
