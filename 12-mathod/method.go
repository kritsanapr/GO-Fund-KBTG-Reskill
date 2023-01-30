package main

import "fmt"

type User struct {
	Username  		string
	FullName  		string
	ProfilePicUrl 	string
}

// info private
// Info public

// Method of type User
func (u User) info() {
	fmt.Printf("Username: %v \n", u.Username)
	fmt.Printf("FullName: %v \n", u.FullName)
	fmt.Printf("ProfilePicUrl: %v \n", u.ProfilePicUrl)
}

// Function of type User
func info(u User) {
	fmt.Printf("Username: %v \n", u.Username)
	fmt.Printf("FullName: %v \n", u.FullName)
	fmt.Printf("ProfilePicUrl: %v \n", u.ProfilePicUrl)
}

func main() {
	u := User {
		Username: "Golang",
		FullName: "Basic Golang",
		ProfilePicUrl: "https://golang.org/doc/gopher/frontpage.png",
	}

	// Function
	info(u)

	// Method of type User
	u.info()
}