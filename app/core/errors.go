package core

import "fmt"

type WithdrawError struct{}

func (w WithdrawError) Error() string {
	return fmt.Sprintf("the requested withdraw amount is bigger than your actual balance")
}
