package main

// Atm represents an ATM machine with its current state and balance.
type Atm struct {
	atmState         AtmState
	twoThousandNotes int
	fiveHundredNotes int
	oneHundredNotes  int
	balance          float64
}

func NewAtm() *Atm {
	// Initializing ATM in IdleState.
	return &Atm{
		atmState: &IdleState{},
	}
}

// SetAtmState changes the state of the ATM.
func (a *Atm) setState(state AtmState) {
	a.atmState = state
}

// SetAtmBalance sets the balance and note counts in the ATM.
func (a *Atm) SetAtmBalance(balance float64, twoThousandNotes int, fiveHundredNotes int, oneHundredNotes int) {
	a.balance = balance
	a.twoThousandNotes = twoThousandNotes
	a.fiveHundredNotes = fiveHundredNotes
	a.oneHundredNotes = oneHundredNotes
}

// getCurrentState returns the current state of the ATM.
func (a *Atm) getCurrentState() AtmState {
	return a.atmState
}
