package main

import "fmt"

type User struct {
	Username string
	FullName string
	ProfilePicUrl string
}

func main() {
	userName := "golang"
	fullName := "Basic Golang"
	profilePicture := "https://golang.org/doc/gopher/frontpage.png"
	fmt.Println(userName, fullName, profilePicture)

	// Initial struct
	// Assign value to struct
	// 1
	// u := User{}
	// u.Username = userName
	// u.FullName = fullName
	// u.ProfilePicUrl = profilePicture

	// 2
	u := User{
		Username: userName,
		FullName: fullName,
		ProfilePicUrl: profilePicture,
	}

	fmt.Printf("%#v \n", u)

	// Read one line from struct
	name := u.Username
	fmt.Printf("Username : %v \n", name)
}