package main

import (
	"fmt"
)

func info(name, message string, age int) {
	fmt.Printf("This is function %v \n", name)
	fmt.Printf("Message: %v \n", message)
	fmt.Printf("Age: %v \n", age)
	fmt.Println("====================================")
}

func today() string {
	return "Today is Sunday"
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	info("Bank", "Hello", 24)
	info("Praew", "Hi", 25)

	today := today()
	fmt.Println(today)

	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}