package main

import (
	"fmt"

	"github.com/kritsanapr/igapp/time"
	"github.com/kritsanapr/igapp/user"
)

func main() {
	user.Info("John", "Hello", 25)

	t := time.Today()
	fmt.Println("Today is: ", t)
}