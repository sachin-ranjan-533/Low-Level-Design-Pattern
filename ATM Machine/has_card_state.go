package main

import "fmt"

type HasCardState struct{}

func (hcs *HasCardState) insertCard(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (hcs *HasCardState) authenticateUser(atm *Atm, card *Card, pin string) {
	if card.pin == pin {
		fmt.Println("Authentication successful. You can now select an operation.")
		atm.setState(&SelectOperationState{})
	} else {
		fmt.Println("Cannot authenticate in idle state. Please insert your card first.")
		atm.setState(&IdleState{})
	}
}

func (hcs *HasCardState) selectOperation(atm *Atm, operation string) {
	fmt.Println("Cannot perform this operation here.")
}

func (hcs *HasCardState) checkBalance(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (hcs *HasCardState) withdrawCash(atm *Atm, card *Card, amount float64) {
	fmt.Println("Cannot perform this operation here.")
}
