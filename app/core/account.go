package core

import "fmt"

type Account struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance" default:"0"`
	UserID  uint64
}

func (a *Account) Deposit(amount float32) {
	a.Balance = a.Balance + amount
}

func (a *Account) Withdraw(amount float32) error {
	if amount > a.Balance {
		return new(WithdrawError)
	}
	a.Balance = a.Balance - amount
	return nil
}

type WithdrawError struct{}

func (w WithdrawError) Error() string {
	return fmt.Sprintf("the requested withdraw amount is bigger than your actual balance")
}
