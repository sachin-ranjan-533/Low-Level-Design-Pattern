package vehicle

import "fmt"

type OrdinaryCar struct{}

func (oc *OrdinaryCar) Drive() {
	fmt.Println("Driving an ordinary car")
}
