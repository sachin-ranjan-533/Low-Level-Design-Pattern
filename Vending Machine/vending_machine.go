package main

type VendingMachine struct {
	State     State
	Inventory Inventory
	Cashs     []Cash
}

func (vm *VendingMachine) NextState(s State) {
	vm.State = s
}

func (vm *VendingMachine) GetState() State {
	return vm.State
}
