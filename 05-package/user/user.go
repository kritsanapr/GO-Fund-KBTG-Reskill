package user

import "fmt"

func Info(name string, msg string, age int) {
	fmt.Println("Name: ", name)
	fmt.Println("Message: ", msg)
	fmt.Println("Age: ", age)
}