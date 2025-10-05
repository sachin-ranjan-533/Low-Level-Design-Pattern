package main

import "fmt"

type TwoThousandWithdrawalProcessor struct {
	nextWithdrawalProcessor WithdrawalProcessor
}

func (ttwp *TwoThousandWithdrawalProcessor) ProcessWithdrawal(atm *Atm, amount float64) bool {
	// Calculate the number of 2000-rupee notes required
	amountAvailable := 2000.00 * float64(atm.twoThousandNotes)

	if amount < 2000 {
		// If the requested amount is less than 2000, forward to the next processor
		return ttwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, amount)
	} else {
		// If the ATM can fulfill the requested amount with available notes
		if amountAvailable >= amount {
			// Calculate how many notes we can dispense
			notesToDispense := int(amount / 2000)
			if notesToDispense > atm.twoThousandNotes {
				notesToDispense = atm.twoThousandNotes // Can't dispense more notes than available
			}

			// Reduce the number of available 2000 notes in the ATM
			atm.twoThousandNotes -= notesToDispense

			// Remaining amount after dispensing the 2000 notes
			remainingAmount := amount - float64(notesToDispense)*2000

			// Print the dispensed notes
			fmt.Printf("Dispensed %d notes of 2000\n", notesToDispense)

			// If there is any remaining amount, pass it to the next processor
			if remainingAmount > 0 && ttwp.nextWithdrawalProcessor != nil {
				return ttwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, remainingAmount)
			}

			return true
		} else {
			// If ATM cannot fully fulfill the requested amount
			remainingAmount := amount - amountAvailable
			// Set the ATM notes to 0 as they are all dispensed
			atm.twoThousandNotes = 0

			// Forward the remaining amount to the next processor
			if ttwp.nextWithdrawalProcessor != nil {
				return ttwp.nextWithdrawalProcessor.ProcessWithdrawal(atm, remainingAmount)
			}

			// Print the dispensed notes
			fmt.Printf("Dispensed %d notes of 2000\n", atm.twoThousandNotes)
			return true
		}
	}
}
