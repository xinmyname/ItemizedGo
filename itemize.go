package main

import "fmt"
import "itemizedgo/model"

func main() {

	item := model.Item{}

	fmt.Printf("Hello %s!\n", item.String())
}
