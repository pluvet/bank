package eventpublisher

type Event interface {
	GetName() string
}

type EventPublisher interface {
	NewEvent(Event) bool
}
