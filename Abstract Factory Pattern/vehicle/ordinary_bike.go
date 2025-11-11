package vehicle

import "fmt"

type OrdinaryBike struct{}

func (ob *OrdinaryBike) Drive() {
	fmt.Println("Driving an ordinary bike")
}
