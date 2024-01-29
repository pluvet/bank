package eventpublisher

import "github.com/pluvet/bank/app/core"

type FakeEventPublisher struct{}

func (e *FakeEventPublisher) NewEvent(event core.Event) bool {
	return true
}
