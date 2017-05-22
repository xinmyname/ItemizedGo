package model

import "sync"

// Descriptor represents a unique identifier for an item
type Descriptor struct {
}

var instance *Descriptor
var once sync.Once

// DefaultDescriptor holds the default descriptor for items
func DefaultDescriptor() *Descriptor {

	once.Do(func() {
		instance = &Descriptor{}
	})

	return instance
}
