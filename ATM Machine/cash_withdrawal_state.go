package main

import "fmt"

type CashWithdrawalState struct{}

func (cws *CashWithdrawalState) insertCard(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (cws *CashWithdrawalState) authenticateUser(atm *Atm, card *Card, pin string) {
	fmt.Println("Cannot perform this operation here.")
}

func (cws *CashWithdrawalState) selectOperation(atm *Atm, operation string) {
	fmt.Println("Cannot perform this operation here.")
}

func (cws *CashWithdrawalState) checkBalance(atm *Atm, card *Card) {
	fmt.Println("Cannot perform this operation here.")
}

func (cws *CashWithdrawalState) withdrawCash(atm *Atm, card *Card, amount float64) {
	// implement cash withdrawal logic
	if amount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive amount.")
		return
	}

	if amount > atm.balance {
		fmt.Println("ATM has insufficient cash.")
		return
	}

	if amount > card.bankAccount.balance {
		fmt.Println("Insufficient amount in your bank account.")
		return
	}

	withdrawalProcessor := &TwoThousandWithdrawalProcessor{nextWithdrawalProcessor: &FiveHundredWithdrawalProcessor{nextWithdrawalProcessor: &OneHundredWithdrawalProcessor{}}}
	if !withdrawalProcessor.ProcessWithdrawal(atm, amount) {
		fmt.Println("Cannot dispense the requested amount with available denominations.")
		return
	}
	fmt.Printf("Please take your cash: %f\n", amount)
	card.bankAccount.debit(amount)
	fmt.Printf("New balance: %f\n", card.bankAccount.balance)
	atm.setState(&IdleState{})
}
