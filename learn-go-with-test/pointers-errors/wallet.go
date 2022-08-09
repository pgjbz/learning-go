package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}


var ErrNoEnoughBalance error = errors.New("no enough balance")


func (w *Wallet) String() string {
	return fmt.Sprintf("%d BTC", w.balance)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) (err error) {
	if w.balance < amount {
		return ErrNoEnoughBalance
	}
	w.balance -= amount
	return
}

func (w *Wallet) Balance() Bitcoin {
	return Bitcoin(w.balance)
}