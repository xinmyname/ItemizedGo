package model

import (
	"fmt"
	"strconv"
)

// Slot contains a quantity and an Item
type Slot struct {
	Quantity int
	Item     Item
}

// MakeSlot creates a new Slot
func MakeSlot(item Item) Slot {
	return Slot{Quantity: 1, Item: item}
}

// String returns a nicely formatted string describing the slot
func (slot Slot) String() string {

	var quantityText string
	var text string

	if slot.Quantity == 1 {
		text = slot.Item.String()
	} else {
		text = fmt.Sprintf("%ss", slot.Item.String())
	}

	switch slot.Quantity {
	case 1:
		quantityText = "An"
	case 0:
		quantityText = "No"
	case 2:
		quantityText = "Two"
	case 3:
		quantityText = "Three"
	case 4:
		quantityText = "Four"
	case 5:
		quantityText = "Five"
	case 6:
		quantityText = "Six"
	case 7:
		quantityText = "Seven"
	case 8:
		quantityText = "Eight"
	case 9:
		quantityText = "Nine"
	default:
		quantityText = strconv.Itoa(slot.Quantity)
	}

	return fmt.Sprintf("%s %s", quantityText, text)
}
