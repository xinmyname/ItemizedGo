package infrastructure

import "itemizedgo/model"

// ItemFactory creates items
type ItemFactory struct {
}

// MakeItem creates a new item
func (ItemFactory) MakeItem() model.Item {
	return model.Item{Descriptor: model.DefaultDescriptor()}
}
