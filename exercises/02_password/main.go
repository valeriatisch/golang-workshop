package main

import (
	"fmt"

	"github.com/valeriatisch/golang-workshop/exercises/02_passwords/pws"
)

func main() {
	// Prompt the user to input a password
	var password string
	fmt.Print("Enter a password: ")
	fmt.Scanln(&password)

	// Validate the password
	isValid := pws.ValidatePassword(password)
	if isValid {
		fmt.Println("Password is valid")
	} else {
		fmt.Println("Password is invalid")
	}
}
