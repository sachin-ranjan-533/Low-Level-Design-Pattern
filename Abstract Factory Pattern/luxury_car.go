package main

import "fmt"

type LuxuryCar struct{}

func (lc *LuxuryCar) drive() {
	fmt.Println("Driving a luxury car")
}
