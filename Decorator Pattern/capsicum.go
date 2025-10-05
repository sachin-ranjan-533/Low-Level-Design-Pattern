package main

type Mushroom struct {
	pizza Pizza
}

func (m *Mushroom) cost() int {
	return m.pizza.cost() + 20
}
