package main

import (
	"fmt"
	"example.com/struct/user"
)

func main() {
	userFirstName := getUserData("Enter your first name:")
	userLastName := getUserData("Enter your last name:")
	userBirthdate := getUserData("Enter your birth date (YYYY-MM-DD):")

	userData, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	userData.ClearingUserData()
	userData.OutputMethodData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
