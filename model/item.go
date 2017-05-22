package model

// Item is an item
type Item struct {
	Descriptor *Descriptor
}

// String returns a description of the item, based on the Descriptor
func (Item) String() string {
	return "item"
}
