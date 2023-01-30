package main

import (
	// "errors"
	"fmt"
	"log"
)
func ReadFile(name string) (string, error) {
	// return "", errors.New("file not found")
	return "", nil

}

func main() {
	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Read file success", data)
}