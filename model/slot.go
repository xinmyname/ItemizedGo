package model

import (
	"fmt"
	"strconv"

	"itemizedgo/infrastructure"
)

// Slot contains a quantity and an Item
type Slot struct {
	Quantity int
	Item     Item
}

// NewSlot creates a new Slot
func NewSlot(item Item) Slot {
	return Slot{Quantity: 1, Item: item}
}

// String returns a nicely formatted string describing the slot
func (slot Slot) String() string {

	var quantityText string
	text := infrastructure.PluralOf(slot.Item.String(), slot.Quantity)

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
