package main

type State interface {
	PressInsertCashButton(vm *VendingMachine)
	InsertCash(vm *VendingMachine, cash Cash)
	CancelButton(vm *VendingMachine)
	SelectProductButton(vm *VendingMachine)
	SelectProduct(vm *VendingMachine, code int)
	ReturnExtraCash(vm *VendingMachine, amount int)
	DispenseProduct(vm *VendingMachine)
}
