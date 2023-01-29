package main

import "fmt"

func main() {
	var day string  = "Monday"
	const month string = "January"

	fmt.Println("This month is", month)

	fmt.Printf("Today is %v \n", day)

	day = "Tuesday"
	fmt.Println("Today is", day)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday", Sunday)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)



}