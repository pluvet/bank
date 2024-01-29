package handlers

import "fmt"

type ErrorEventIsNotSupported struct{}

func (e ErrorEventIsNotSupported) Error() string {
	return fmt.Sprintf("the event received in handler is not one of the event types supported")
}
