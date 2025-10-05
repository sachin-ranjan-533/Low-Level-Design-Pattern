package main

import "fmt"

type Car struct{}

func (c *Car) drive() {
	fmt.Println("Driving a Car")
}
