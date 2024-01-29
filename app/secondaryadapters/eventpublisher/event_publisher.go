package eventpublisher

import (
	"fmt"
	"sync"

	"github.com/pluvet/bank/app/core"
)

type Handler interface {
	HandleEvent(core.Event) error
}

type EventPublisher struct {
	handlers map[string][]Handler
}

func NewEventPublisher(handlers map[string][]Handler) *EventPublisher {
	var eventPublisher = new(EventPublisher)
	eventPublisher.handlers = handlers
	return eventPublisher
}

func (e *EventPublisher) NewEvent(event core.Event) bool {
	eventWasPublished := make(chan bool)
	go e.processEvent(event, eventWasPublished)
	return <-eventWasPublished
}

func (e *EventPublisher) processEvent(event core.Event, eventWasPublished chan bool) {
	var wg sync.WaitGroup
	eventHandlers := e.handlers[event.GetName()]
	for i := range eventHandlers {
		var handler = eventHandlers[i]
		go e.handleEvent(handler, event, &wg)
	}
	eventWasPublished <- true
	wg.Wait()
}

func (e *EventPublisher) handleEvent(handler Handler, event core.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	err := handler.HandleEvent(event)
	fmt.Println(err)
}
