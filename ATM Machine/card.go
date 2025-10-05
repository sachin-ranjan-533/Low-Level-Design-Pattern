package main

type Card struct {
	number      string
	pin         string
	bankAccount *BankAccount
}

func NewCard(number string, bankAccount *BankAccount) *Card {
	return &Card{
		number:      number,
		pin:         "1234", // Default PIN for simplicity
		bankAccount: bankAccount,
	}
}
