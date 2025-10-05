package main

type AtmState interface {
	insertCard(atm *Atm, card *Card)
	authenticateUser(atm *Atm, card *Card, pin string)
	selectOperation(atm *Atm, operation string)
	checkBalance(atm *Atm, card *Card)
	withdrawCash(atm *Atm, card *Card, amount float64)
}
