package main

import (
	"fmt"
	"itemizedgo/model"
	"os"
	"strconv"
)

func main() {

	count := 1

	if len(os.Args) >= 2 {

		countArg, err := strconv.Atoi(os.Args[1])

		if err == nil {
			count = countArg
		}
	}

	inventory := model.NewInventory()

	for i := 0; i < count; i++ {
		inventory.AddItem(model.NewItem())
	}

	fmt.Println("You have:")

	for _, slot := range inventory.Slots() {
		fmt.Printf("  %s\n", slot.String())
	}
}
