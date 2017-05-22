package model

// Inventory is a set of slots for Items
type Inventory struct {
	slots map[Descriptor]*Slot
}

// MakeInventory creates an Inventory object
func MakeInventory() Inventory {
	return Inventory{slots: make(map[Descriptor]*Slot)}
}

// AddItem adds an item to an existing slot, or creates a new slot
func (i Inventory) AddItem(item Item) {

	if slot, ok := i.slots[*item.Descriptor]; ok {
		(*slot).Quantity++
	} else {
		newSlot := MakeSlot(item)
		i.slots[*item.Descriptor] = &newSlot
	}
}

// Slots returns an array of all of the slots in the inventory
func (i Inventory) Slots() []Slot {

	slots := make([]Slot, 0, len(i.slots))

	for _, value := range i.slots {
		slots = append(slots, *value)
	}

	return slots
}
