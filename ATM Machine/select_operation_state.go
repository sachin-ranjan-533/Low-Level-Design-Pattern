package main

import "fmt"

type SelectOperationState struct {
}

func (sos *SelectOperationState) insertCard(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (sos *SelectOperationState) authenticateUser(atm *Atm, card *Card, pin string) {
	fmt.Println("Cannot perform this operation here.")
}

func (sos *SelectOperationState) selectOperation(atm *Atm, operation string) {
	switch operation {
	case "check_balance":
		fmt.Println("Transitioning to Check Balance State.")
		atm.setState(&CheckBalanceState{})
	case "withdraw_cash":
		fmt.Println("Transitioning to Cash Withdrawal State.")
		atm.setState(&CashWithdrawalState{})
	default:
		fmt.Println("Invalid operation selected. Please choose a valid operation.")
	}
}

func (sos *SelectOperationState) checkBalance(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (sos *SelectOperationState) withdrawCash(atm *Atm, card *Card, amount float64) {
	fmt.Println("Cannot perform this operation here.")
}
