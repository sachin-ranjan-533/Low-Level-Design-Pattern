package main

type Inventory struct {
	ItemShelfs []ItemShelf
}

func (i *Inventory) AddItemShelf(itemShelf ItemShelf) {
	i.ItemShelfs = append(i.ItemShelfs, itemShelf)
}
