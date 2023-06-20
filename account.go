package main

import "errors"

type Account struct {
	balance float64
}

func NewAccount() *Account {
	return &Account{
		balance: 0.0,
	}
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	newBalance := acc.balance - amount
	if newBalance < 0.0 {
		return errors.New("insufficient funds")
	} else {
	acc.balance = newBalance
	return nil
	}
} 