package main

import "fmt"

type NormalDrive struct{}

func (nd *NormalDrive) drive() {
	fmt.Println("Driving in normal mode.")
}
