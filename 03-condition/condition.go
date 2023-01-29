package main

import (
	"fmt"
	"time"
)

func main() {
	num := 0
	// If else condition
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else if num%3 == 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is ...")
	}

	// Switch case condition
	today := time.Saturday
	switch today {
		case time.Monday:
			fmt.Println("Today is Monday")
		case time.Tuesday:
			fmt.Println("Today is Tuesday")
		case time.Saturday, time.Sunday:
			fmt.Println("Today is Weekend")
			fallthrough
		default:
			fmt.Println("Today is Me hang day")
	}

}