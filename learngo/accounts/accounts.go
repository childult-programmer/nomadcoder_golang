package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(owner string) *Account {
	newAccount := Account{owner: owner, balance: 0}

	return &newAccount
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	var errNoMoney = errors.New("The amount you are trying to withdraw is more than you have")

	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Balance of the account
func (a Account) Balance() int {
	return a.balance
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. has: ", a.Balance())
}
