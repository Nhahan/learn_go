package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

// NewAccount Constructor function
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Balance method
func (a Account) Balance() int {
	return a.balance
}

// Deposit method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw method
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

var errNoMoney = errors.New("Can't withdraw")
