package main

type BasePizza struct{}

func (bp *BasePizza) cost() int {
	return 100
}

func NewPizza() *BasePizza {
	return &BasePizza{}
}
