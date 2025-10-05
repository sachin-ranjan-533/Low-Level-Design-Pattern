package main

func main () {
	basePizza := NewPizza()
	mushroomPizza := &Mushroom{pizza: basePizza}
	extraCheeseMushroomPizza := &ExtraCheese{pizza: mushroomPizza}
	println("Cost of base pizza:", basePizza.cost())
	println("Cost of mushroom pizza:", mushroomPizza.cost())
	println("Cost of extra cheese mushroom pizza:", extraCheeseMushroomPizza.cost())
}