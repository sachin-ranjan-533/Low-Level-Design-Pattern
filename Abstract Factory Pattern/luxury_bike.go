package main

import "fmt"

type LuxuryBike struct{}

func (lb *LuxuryBike) drive() {
	fmt.Println("Driving a luxury bike")
}
