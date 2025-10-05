package main

import "fmt"

type RacingDrive struct{}

func (rd *RacingDrive) drive() {
	fmt.Println("Driving in racing mode.")
}
