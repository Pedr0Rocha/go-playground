package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	wallet := Wallet{Bitcoin(100)}

	fmt.Println(wallet.Balance())

	wallet.Deposit(Bitcoin(10))

	fmt.Println(wallet.Balance())

	wallet.Withdraw(Bitcoin(200))

	fmt.Println(wallet.Balance())

	wallet.Withdraw(Bitcoin(60))

	fmt.Println(wallet.Balance())
}
