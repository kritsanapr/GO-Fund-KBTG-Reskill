package main

import "fmt"

 func main() {
	langs := []string{"golang", "python", "java", "c++", "c", "c#"}
	fmt.Printf("langs: %#v \n", langs)

	n := langs[0]
	fmt.Printf("n: %#v\n", n)

	langs[1] = "Pythonistas"
	fmt.Printf("n: %#v\n", langs)

	l := len(langs)
	fmt.Printf("len: %#v\n", l)

	langs = append(langs, "ruby", "F#")
	fmt.Printf("Append: %#v\n", langs)

 }