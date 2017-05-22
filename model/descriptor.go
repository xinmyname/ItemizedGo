package model

import "sync"
import "fmt"

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

	fmt.Printf("D: %p\n", instance)

	return instance
}
