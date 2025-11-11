package vehicle

import "fmt"

type LuxuryBike struct{}

func (lb *LuxuryBike) Drive() {
	fmt.Println("Driving an luxury bike")
}
