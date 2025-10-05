package main

func main() {
	// Initialize an ATM room and print it
	atmRoom := NewAtmRoom()
	atmRoom.initializeAtmRoom()

	atmRoom.atm.getCurrentState().insertCard(atmRoom.atm, atmRoom.user.card)
	atmRoom.atm.getCurrentState().authenticateUser(atmRoom.atm, atmRoom.user.card, "1234")
	atmRoom.atm.getCurrentState().selectOperation(atmRoom.atm, "withdraw_cash")
	atmRoom.atm.getCurrentState().withdrawCash(atmRoom.atm, atmRoom.user.card, 2300)
	// atmRoom.atm.getCurrentState().selectOperation(atmRoom.atm, "check_balance")
	// atmRoom.atm.getCurrentState().checkBalance(atmRoom.atm, atmRoom.user.card)
}
