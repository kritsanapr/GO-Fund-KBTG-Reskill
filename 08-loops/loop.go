package main

import "fmt"

func main() {
	// Normal loop
	for i := 0; i<10; i++ {
		fmt.Printf("Looping : %v \n", i)
	}

	fmt.Println("--------------")

	// while loop
	y := 0
	for y<10 {
		fmt.Printf("While loop :%v \n", y)
		y++
	}

	fmt.Println("-------Loop with index-------")

	langs := []string{"golang", "python", "java", "c++", "c", "c#"}
	for index, lang := range langs {
		fmt.Printf("Index : %v, Value : %v \n", index, lang)
	}

	fmt.Println("--------Loop array without index------")
	for _, lang := range langs {
		fmt.Printf("Value : %v \n",  lang)
	}

	fmt.Println("--------Loop array using index only------")
	for i := range langs {
		fmt.Printf("Index : %v \n",  i)
	}


}