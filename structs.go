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

func (u *user) outputMethodData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearingUserData() {
	fmt.Println("Clearing user data...")
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.createdAt = time.Time{}
	fmt.Println("User data cleared.")
}

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Time{},
	}
}

func main() {
	userFirstName := getUserData("Enter your first name:")
	userLastName := getUserData("Enter your last name:")
	userBirthdate := getUserData("Enter your birth date (YYYY-MM-DD):")

	userData := newUser(userFirstName, userLastName, userBirthdate)

	outputUserData(userData)
	userData.clearingUserData()
	userData.outputMethodData()
}

func outputUserData(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
