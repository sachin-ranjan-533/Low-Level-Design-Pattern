package main

type Number struct {
	value int
}

func (n *Number) Evaluate() int {
	return n.value
}
