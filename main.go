package main

import (
	"fmt"
	"net/http"
)

type BankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

var account = NewAccount()

func Deposit(w http.ResponseWriter, r *http.Request) {
	account.Deposit(75.0)
	currentBalance := account.GetBalance()
	fmt.Fprintf(w, "Balance: $%.2f\n", currentBalance)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	if err := account.Withdraw(5.35); err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
	}
	currentBalance := account.GetBalance()
	fmt.Fprintf(w, "Balance: $%.2f\n", currentBalance)
}

func Balance(w http.ResponseWriter, r *http.Request) {
	currentBalance := account.GetBalance()
	if currentBalance < 25 {
		fmt.Fprintf(w, "%s\n", "Low ")
	}
	fmt.Fprintf(w, "Balance: $%.2f\n", currentBalance)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/deposit", Deposit)
	http.HandleFunc("/withdraw", Withdraw)
	http.HandleFunc("/balance", Balance)
	http.ListenAndServe(":8090", nil)
}