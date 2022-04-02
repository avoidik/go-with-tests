package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Deposit(howMany Bitcoin) {
	w.balance += howMany
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(howMany Bitcoin) error {
	if howMany > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= howMany
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
