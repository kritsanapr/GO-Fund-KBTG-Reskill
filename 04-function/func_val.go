package main

import "fmt"

var name string = "Bank"


// func plus (a int, b int) int {
	// 	return a + b
	// }

func main() {
	var name  string = "Bank"
	fmt.Println(name)
	fmt.Printf("Type: %T \n", name)

	plus := func (a int, b int) int { return a + b }
	p := plus(1, 2)
	fmt.Println(p)
	fmt.Printf("Type: %T \n", plus)
}