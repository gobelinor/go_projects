package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// define what "%s" Bitcoin looks like
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Bitcoin is now part of the Stringer interface because it implements String()
type Stringer interface {
	String() string
}

//In Go, when you call a function or a method the arguments are copied.

func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
}

// *Wallet :  a pointer to a Wallet / struct pointer
func (w *Wallet) Balance() Bitcoin {
	// No need to dereference (*w).balance, it's automatic
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(val Bitcoin) error {
	if val > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= val
	return nil
}
