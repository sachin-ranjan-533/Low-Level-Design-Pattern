package main

// Define the AtmRoom struct
type AtmRoom struct {
	user *User
	atm  *Atm
}

// NewAtmRoom function creates a new AtmRoom instance
func NewAtmRoom() *AtmRoom {
	return &AtmRoom{}
}

// initializeAtmRoom sets up the ATM room with a user and an ATM
func (ar *AtmRoom) initializeAtmRoom() {
	bankAccount := NewBankAccount(3000)
	card := NewCard("1234567890123456", bankAccount)
	user := NewUser("Sachin Ranjan", card, bankAccount)

	atm := NewAtm()
	atm.SetAtmBalance(20000, 5, 10, 20)
	ar.atm = atm
	ar.user = user
}
