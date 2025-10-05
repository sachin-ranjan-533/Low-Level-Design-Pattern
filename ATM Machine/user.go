package main

type User struct {
	userName    string
	card        *Card
	bankAccount *BankAccount
}

func NewUser(userName string, card *Card, bankAccount *BankAccount) *User {
	return &User{
		userName:    userName,
		card:        card,
		bankAccount: bankAccount,
	}
}
