package main

import "fmt"

/*
	basic types
	sting
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte - alias for uint8

	rune //alias for int32
		represents a Unicode code point

	float32 float64

	complex64 complex128
*/

var name string = "Gopher"

func main() {
	// var firstName = "Kritsana"
	// var lastname string = "Kritsana"
	// name_local := "Bank"
	// var name string

	var name string
	s := "Gopher"
	i := 77
	f := 3.14
	b := true
	r := 'a'


	fmt.Printf("Hello, World! %v", name)

	fmt.Printf("Hello %s !!!\n", s)
	fmt.Printf("Hello %v World! %v \n", i, s)
	fmt.Printf("Hello %v World! %v \n", f, s)
	fmt.Printf("Hello %v World! %v \n", b, s)
	fmt.Printf("Hello %v World! %v \n", r, s)
}