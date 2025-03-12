package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	asserteq := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if err != want {
			t.Errorf("got '%q', want '%q'", err, want)
		}
	}
	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		//In Go, when you call a function or a method the arguments are copied.
		wallet.Deposit(Bitcoin(10))
		asserteq(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		asserteq(t, wallet, Bitcoin(5))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		asserteq(t, wallet, startingBalance)
	})
}
