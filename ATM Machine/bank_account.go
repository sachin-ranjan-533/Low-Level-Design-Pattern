package main

type BankAccount struct {
	balance float64
}

func NewBankAccount(balance float64) *BankAccount {
	return &BankAccount{
		balance: balance,
	}
}

func (ba *BankAccount) debit(amount float64) {
	ba.balance -= amount
}

func (ba *BankAccount) credit(amount float64) {
	ba.balance += amount
}
