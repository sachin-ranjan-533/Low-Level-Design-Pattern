package main

import "fmt"

type IdleState struct{}

func (is *IdleState) insertCard(atm *Atm, card *Card) {
	fmt.Println("Card inserted. Please authenticate.")
	atm.setState(&HasCardState{})
}

func (is *IdleState) authenticateUser(atm *Atm, card *Card, pin string) {
	fmt.Println("Cannot perform this operation here.")
}

func (is *IdleState) selectOperation(atm *Atm, operation string) {
	fmt.Println("Cannot perform this operation here.")
}

func (is *IdleState) checkBalance(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (is *IdleState) withdrawCash(atm *Atm, card *Card, amount float64) {
	fmt.Println("Cannot perform this operation here.")
}
