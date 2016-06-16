package main

import (
	"fmt"
)

func main() {
	fmt.Println("Importing ship data")
	importShips()
	fmt.Println("Done")
	fmt.Println("Importing item data")
	importItems()
	fmt.Println("Done")
}
