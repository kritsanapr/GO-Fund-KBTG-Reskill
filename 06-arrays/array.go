package main

import "fmt"

func showAll(langs []string) {
		fmt.Printf("value: %#v \n", langs)
}

func main() {
	var langs = []string{"golang", "python", "ruby", "javascript"}
	fmt.Printf("langes: %#v \n", langs)

	n := langs[2]
	fmt.Printf("n: %#v \n", n)

	langs[0] = "java"
	fmt.Printf("langes: %#v \n", langs)

	langs = append(langs, "c++")
	showAll(langs)
}