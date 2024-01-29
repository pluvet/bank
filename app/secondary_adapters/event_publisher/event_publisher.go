package secondary_adapters

import (
	"fmt"
	"sync"
)

type Event interface {
	GetName() string
}

type Handler interface {
	HandleEvent(Event) error
}

type EventPublisher struct {
	handlers map[string][]Handler
}

func NewEventPublisher(handlers map[string][]Handler) *EventPublisher {
	var eventPublisher = new(EventPublisher)
	eventPublisher.handlers = handlers
	return eventPublisher
}

func (e *EventPublisher) NewEvent(event Event) bool {
	eventWasPublished := make(chan bool)
	go e.processEvent(event, eventWasPublished)
	return <-eventWasPublished
}

func (e *EventPublisher) processEvent(event Event, eventWasPublished chan bool) {
	var wg sync.WaitGroup
	eventHandlers := e.handlers[event.GetName()]
	for i := range eventHandlers {
		var handler = eventHandlers[i]
		go e.handleEvent(handler, event, &wg)
	}
	eventWasPublished <- true
	wg.Wait()
}

func (e *EventPublisher) handleEvent(handler Handler, event Event, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	err := handler.HandleEvent(event)
	fmt.Println(err)
}
