package main

import (
	"errors"
	"fmt"
)

type IdleState struct{}

func (is *IdleState) PressInsertCashButton(vm *VendingMachine) {
	fmt.Println("Insert cash button pressed. Switching to HasMoneyState.")
	vm.NextState(&HasMoneyState{})
}

func (is *IdleState) InsertCash(vm *VendingMachine, cash Cash) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (is *IdleState) CancelButton(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (is *IdleState) SelectProductButton(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (is *IdleState) SelectProduct(vm *VendingMachine, code int) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (is *IdleState) ReturnExtraCash(vm *VendingMachine, amount int) {
	fmt.Println(errors.New("This operation is not allowed!"))
}

func (is *IdleState) DispenseProduct(vm *VendingMachine) {
	fmt.Println(errors.New("This operation is not allowed!"))
}
