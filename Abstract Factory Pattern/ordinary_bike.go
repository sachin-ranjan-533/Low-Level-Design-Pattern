package main

import "fmt"

type OrdinaryBike struct{}

func (ob *OrdinaryBike) drive() {
	fmt.Println("Driving a ordinary bike")
}
