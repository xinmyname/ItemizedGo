package model

// Item is an item
type Item struct {
	Descriptor *Descriptor
}

// NewItem creates a new item
func NewItem() Item {
	return Item{Descriptor: DefaultDescriptor()}
}

// String returns a description of the item, based on the Descriptor
func (Item) String() string {
	return "item"
}
