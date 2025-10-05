package main

import (
	"errors"
	"fmt"
)

type DispenseState struct{}

func (ds *DispenseState) PressInsertCashButton(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (ds *DispenseState) InsertCash(vm *VendingMachine, cash Cash) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (ds *DispenseState) CancelButton(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (ds *DispenseState) SelectProductButton(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (ds *DispenseState) SelectProduct(vm *VendingMachine, code int) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (ds *DispenseState) ReturnExtraCash(vm *VendingMachine, amount int) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (ds *DispenseState) DispenseProduct(vm *VendingMachine) {
	fmt.Println("Dispensing product...")

	// Reset inserted cash after product is dispensed
	vm.Cashs = nil

	fmt.Println("Product dispensed. Returning to IdleState.")
	vm.NextState(&IdleState{})
}
