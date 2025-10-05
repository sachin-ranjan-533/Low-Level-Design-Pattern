package main

import "fmt"

type CheckBalanceState struct{}

func (cbs *CheckBalanceState) insertCard(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (cbs *CheckBalanceState) authenticateUser(atm *Atm, card *Card, pin string) {
	fmt.Println("Cannot perform this operation here.")
}

func (cbs *CheckBalanceState) selectOperation(atm *Atm, operation string) {
	fmt.Println("Cannot perform this operation here.")
}

func (cbs *CheckBalanceState) checkBalance(atm *Atm, card *Card) {
	fmt.Printf("Your current balance is: %f", card.bankAccount.balance)
	atm.setState(&IdleState{})
}

func (cbs *CheckBalanceState) withdrawCash(atm *Atm, card *Card, amount float64) {
	fmt.Println("Cannot perform this operation here.")
}
