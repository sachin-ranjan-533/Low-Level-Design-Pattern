package main

import "fmt"

type FiveHundredWithdrawalProcessor struct {
	nextWithdrawalProcessor WithdrawalProcessor
}

func (fhwp *FiveHundredWithdrawalProcessor) ProcessWithdrawal(atm *Atm, amount float64) bool {
	// Calculate the number of 500-rupee notes required
	amountAvailable := 500.00 * float64(atm.fiveHundredNotes)

	if amount < 500 {
		// If the requested amount is less than 500, forward to the next processor
		return fhwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, amount)
	} else {
		// If the ATM can fulfill the requested amount with available notes
		if amountAvailable >= amount {
			// Calculate how many notes we can dispense
			notesToDispense := int(amount / 500)
			if notesToDispense > atm.fiveHundredNotes {
				notesToDispense = atm.fiveHundredNotes // Can't dispense more notes than available
			}

			// Reduce the number of available 500 notes in the ATM
			atm.fiveHundredNotes -= notesToDispense

			// Remaining amount after dispensing the 500 notes
			remainingAmount := amount - float64(notesToDispense)*500

			// Print the dispensed notes
			fmt.Printf("Dispensed %d notes of 500\n", notesToDispense)

			// If there is any remaining amount, pass it to the next processor
			if remainingAmount > 0 && fhwp.nextWithdrawalProcessor != nil {
				return fhwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, remainingAmount)
			}

			return true
		} else {
			// If ATM cannot fully fulfill the requested amount
			remainingAmount := amount - amountAvailable
			// Set the ATM notes to 0 as they are all dispensed
			atm.fiveHundredNotes = 0

			// Print the dispensed notes
			fmt.Printf("Dispensed %d notes of 500\n", atm.fiveHundredNotes)

			// Forward the remaining amount to the next processor
			if fhwp.nextWithdrawalProcessor != nil {
				return fhwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, remainingAmount)
			}

			return true
		}
	}
}
