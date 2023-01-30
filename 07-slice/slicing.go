package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "java", "c++", "c", "c#"}
	fmt.Printf("langs: %#v \n", langs)

	a := langs[0:2]
	fmt.Printf("Slicing A : %#v \n", a)

	b := langs[1:3]
	fmt.Printf("Slicing B : %#v \n", b)

	// c := langs[:len(langs)]
	c := langs[:]
	printSlice(c)
	// fmt.Printf("All langs : %#v \n", c)
}

func printSlice(ns []string) {
	fmt.Printf("printSlice: %#v", ns)
}