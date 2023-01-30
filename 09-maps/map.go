package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		404: "Not Found",
		500: "Internal Server Error",
		// 66: "",
	}

	fmt.Printf("Status : %v \n\n", status)

	l := len(status)
	fmt.Printf("Len Status : %v \n", l)

	status[200] = "OK - Updated"
	fmt.Printf("Update Status : %v \n\n", status)

	// Add new value
	status[201] = "Created"
	fmt.Printf("Add Status : %v \n\n", status)
	fmt.Printf("Len Status : %v \n", len(status))

	value := status[200]
	fmt.Printf("Value Status : %v \n\n", value)
	fmt.Printf("Len Status : %v \n", len(status))

	// Delete value in map
	delete(status, 201)
	fmt.Printf("Delete Status : %v \n\n", status)
	fmt.Printf("Len Status : %v \n", len(status))

	// Select zero value in map
	value = status[202]
	fmt.Printf("Select zero value in map : %v \n", value)


	if v, ok := status[66]; ok { // Short if conditions
		fmt.Printf("Value Status : %v \n", v)
	} else {
		fmt.Printf("Value Status : %v \n", "Not Found")
	}

	// var m map[string]string
	// var m map[string]string = map[string]string{}
	// var m map[string]string = make(map[string]string)
	// var m = make(map[string]string)
	// var m = make(map[string]string)
	m  := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	fmt.Printf("Value m : %v \n", m)
	if m == nil {
		fmt.Printf("Value m is nil \n")
	} else {
		fmt.Printf("Value m is not nil \n")
	}

}