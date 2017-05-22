package main

import (
	"fmt"
	"itemizedgo/infrastructure"
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

	for i := 0; i < count; i++ {

		item := itemFactory.MakeItem()
		fmt.Printf("Hello %s!\n", item.String())
	}

}
