package main

import "fmt"

// Concrete Observer: Customer
type Customer struct {
	id         int
	observable Observable
}

// Called when observable state changes
func (c *Customer) update() {
	fmt.Printf("Available stock count for iPhone is %d. You can place the order.\n", c.observable.getData())
}

// Returns the unique ID of the customer
func (c *Customer) getId() int {
	return c.id
}
