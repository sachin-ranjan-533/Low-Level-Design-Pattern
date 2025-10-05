package main

import "fmt"

type Bike struct{}

func (b *Bike) drive() {
	fmt.Println("Driving a Bike")
}
