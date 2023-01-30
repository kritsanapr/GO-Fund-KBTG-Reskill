package main

import "fmt"

// Zero value of slice is nil or null
func main() {
	// var langs []string // Slice is nil
	langs := []string{} // Slice is not nill
	fmt.Printf("%v\n", langs)

	if langs == nil {
		fmt.Printf(`langs is nil`)
	} else {
		fmt.Printf(`langs is not nil`)
	}
}