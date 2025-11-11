package vehicle

import "fmt"

type LuxuryCar struct{}

func (lc *LuxuryCar) Drive() {
	fmt.Println("Driving an luxury car")
}
