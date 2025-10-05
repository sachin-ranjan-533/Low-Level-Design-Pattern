package main

import (
	"fmt"
)

type SelectionState struct{}

func (ss *SelectionState) PressInsertCashButton(vm *VendingMachine) {
	fmt.Println("This operation is not allowed!")
}

func (ss *SelectionState) InsertCash(vm *VendingMachine, cash Cash) {
	fmt.Println("This operation is not allowed!")
}

func (ss *SelectionState) CancelButton(vm *VendingMachine) {
	fmt.Println("Cancelling transaction. Switching to IdleState.")
	vm.NextState(&IdleState{})
}

func (ss *SelectionState) SelectProductButton(vm *VendingMachine) {
	fmt.Println("This operation is not allowed!")
}

func (ss *SelectionState) SelectProduct(vm *VendingMachine, code int) {
	totalCash := 0
	for _, c := range vm.Cashs {
		totalCash += c.Value
	}

	var selectedItem *ItemShelf
	for i := range vm.Inventory.ItemShelfs {
		if vm.Inventory.ItemShelfs[i].Code == code {
			selectedItem = &vm.Inventory.ItemShelfs[i]
			break
		}
	}

	if selectedItem == nil {
		fmt.Println("Invalid product code.")
		vm.NextState(&IdleState{})
		return
	}

	itemCost := selectedItem.Item.Cost
	if totalCash >= itemCost {
		if totalCash > itemCost {
			ss.ReturnExtraCash(vm, totalCash-itemCost)
		}
		fmt.Println("Selected product. Switching to DispenseState.")
		vm.NextState(&DispenseState{})
	} else {
		fmt.Println("Insufficient funds. Returning to IdleState.")
		vm.NextState(&IdleState{})
	}
}

func (ss *SelectionState) ReturnExtraCash(vm *VendingMachine, amount int) {
	fmt.Printf("Returning extra cash: %d\n", amount)
}

func (ss *SelectionState) DispenseProduct(vm *VendingMachine) {
	fmt.Println("This operation is not allowed!")
}
