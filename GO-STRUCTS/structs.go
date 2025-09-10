package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser, error = user.New(userFirstName, userLastName, userBirthdate)

	if error != nil {
		fmt.Println(error)
		return
	}
	// ... do something awesome with that gathered data!

	admin := user.NewAdmin("email.com", "password")

	admin.OutputUserDetails()
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
