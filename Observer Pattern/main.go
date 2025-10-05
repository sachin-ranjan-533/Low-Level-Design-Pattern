package main

func main() {
	// Create the observable (iPhone stock)
	iPhone := &iPhoneObservable{}

	// Create observers (customers)
	c1 := &Customer{id: 1, observable: iPhone}
	c2 := &Customer{id: 2, observable: iPhone}

	// Register observers with the observable
	iPhone.add(c1)
	iPhone.add(c2)

	// Update stock; all observers will be notified
	iPhone.setData(10)
}
