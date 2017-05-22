package main

import (
	"fmt"
	"itemizedgo/infrastructure"
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

	itemFactory := infrastructure.ItemFactory{}
	inventory := model.MakeInventory()

	for i := 0; i < count; i++ {
		inventory.AddItem(itemFactory.MakeItem())
	}

	fmt.Println("You have:")

	for _, slot := range inventory.Slots() {
		fmt.Printf("  %s\n", slot.String())
	}
}
