package main

func main() {
	idle := &IdleState{}
	vm := &VendingMachine{State: idle}

	// Add items to the inventory
	itemShelf1 := ItemShelf{Code: 1, IsSold: false, Item: Item{ItemType: "SODA", Cost: 10}}
	itemShelf2 := ItemShelf{Code: 2, IsSold: false, Item: Item{ItemType: "COFFEE", Cost: 20}}

	vm.Inventory.AddItemShelf(itemShelf1)
	vm.Inventory.AddItemShelf(itemShelf2)

	// Simulate user interactions
	vm.GetState().PressInsertCashButton(vm)
	vm.GetState().InsertCash(vm, Cash{Value: 1})
	vm.GetState().InsertCash(vm, Cash{Value: 20})
	vm.GetState().SelectProductButton(vm)
	vm.GetState().SelectProduct(vm, 1)
	vm.GetState().DispenseProduct(vm)
}
