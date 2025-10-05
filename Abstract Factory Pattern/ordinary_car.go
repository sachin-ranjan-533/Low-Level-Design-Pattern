package main

import "fmt"

type OrdinaryCar struct{}

func (oc *OrdinaryCar) drive() {
	fmt.Println("Driving a ordinary car")
}
