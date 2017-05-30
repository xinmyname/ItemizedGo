package infrastructure

import (
	"fmt"
)

// PluralOf returns the plural of the specified word if the count is not equal to 1
func PluralOf(word string, count int) string {

	if count == 1 {
		return word
	}

	return fmt.Sprintf("%ss", word)
}
