package eventpublisher

import "github.com/pluvet/bank/app/core"

type EventPublisher interface {
	NewEvent(core.Event) bool
}
