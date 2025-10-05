package main

import "fmt"

type OneHundredWithdrawalProcessor struct {
	nextWithdrawalProcessor WithdrawalProcessor
}

func (ohwp *OneHundredWithdrawalProcessor) ProcessWithdrawal(atm *Atm, amount float64) bool {
	// Calculate the number of 100-rupee notes required
	amountAvailable := 100.00 * float64(atm.oneHundredNotes)

	if amount < 100 {
		// If the requested amount is less than 100, forward to the next processor
		return ohwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, amount)
	} else {
		// If the ATM can fulfill the requested amount with available notes
		if amountAvailable >= amount {
			// Calculate how many notes we can dispense
			notesToDispense := int(amount / 100)
			if notesToDispense > atm.oneHundredNotes {
				notesToDispense = atm.oneHundredNotes // Can't dispense more notes than available
			}

			// Reduce the number of available 100 notes in the ATM
			atm.oneHundredNotes -= notesToDispense

			// Remaining amount after dispensing the 100 notes
			remainingAmount := amount - float64(notesToDispense)*100

			// If there is any remaining amount, pass it to the next processor
			if remainingAmount > 0 && ohwp.nextWithdrawalProcessor != nil {
				return ohwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, remainingAmount)
			}

			// Print the dispensed notes
			fmt.Printf("Dispensed %d notes of 100\n", notesToDispense)
			return true
		} else {
			// If ATM cannot fully fulfill the requested amount
			remainingAmount := amount - amountAvailable
			// Set the ATM notes to 0 as they are all dispensed
			atm.oneHundredNotes = 0

			// Print the dispensed notes
			fmt.Printf("Dispensed %d notes of 100\n", atm.oneHundredNotes)

			// Forward the remaining amount to the next processor
			if ohwp.nextWithdrawalProcessor != nil {
				return ohwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, remainingAmount)
			}

			return true
		}
	}
}
