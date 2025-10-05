package main

type ExtraCheese struct {
	pizza Pizza
}

func (ec *ExtraCheese) cost() int {
	return ec.pizza.cost() + 50
}
