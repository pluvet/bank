package core

type Account struct {
	ID      int     `json:"id" gorm:"primary_key"`
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
