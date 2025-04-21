package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Enter your first name:")
	userLastName := getUserData("Enter your last name:")
	userBirthdate := getUserData("Enter your birth date (YYYY-MM-DD):")

	userData := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserData(&userData)
}

func outputUserData(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
