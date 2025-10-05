package main

import (
	"errors"
	"fmt"
)

type HasMoneyState struct{}

func (hs *HasMoneyState) PressInsertCashButton(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (hs *HasMoneyState) InsertCash(vm *VendingMachine, cash Cash) {
	vm.Cashs = append(vm.Cashs, cash)
	fmt.Printf("Inserting cash %d.\n", cash.Value)
}

func (hs *HasMoneyState) CancelButton(vm *VendingMachine) {
	fmt.Println("Cancelling transaction. Switching to IdleState.")
	vm.NextState(&IdleState{})
}

func (hs *HasMoneyState) SelectProductButton(vm *VendingMachine) {
	fmt.Println("Select product button pressed. Switching to SelectionState.")
	vm.NextState(&SelectionState{})
}

func (hs *HasMoneyState) SelectProduct(vm *VendingMachine, code int) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (hs *HasMoneyState) ReturnExtraCash(vm *VendingMachine, amount int) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (hs *HasMoneyState) DispenseProduct(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}
