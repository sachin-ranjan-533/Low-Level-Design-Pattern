package main

type WithdrawalProcessor interface {
	ProcessWithdrawal(atm *Atm, amount float64) bool
}
