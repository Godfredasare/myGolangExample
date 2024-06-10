package main

import (
	"fmt"

	"github.com/Godfredasare/myGolangExample.git/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthday := getUserData("Please enter your birthday (MM/DD/YY): ")

	user, err := user.New(firstName, lastName, birthday)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	user.PrintOutput()
	user.ClearUser()
	user.PrintOutput()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
