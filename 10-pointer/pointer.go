package main

import "fmt"

func main() {
	me := "Gopher"
	fmt.Printf("me : %v \n", me)

	// Memory address of me
	var addr *string = &me
	// addr := &me
	fmt.Printf("addr : %v \n", addr)
	fmt.Printf("Type : %T \n", addr)

	// Dereference find address value and change value
	*addr = "Penguin"
	fmt.Printf("addr : %v \n", *addr)
}