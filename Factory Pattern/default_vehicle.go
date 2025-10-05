package main

import "fmt"

type DefaultVehicle struct{}

func (dv *DefaultVehicle) drive() {
	fmt.Println("Driving a DefaultVehicle")
}
