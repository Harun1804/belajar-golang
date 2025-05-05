package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) OutputMethodData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearingUserData() {
	fmt.Println("Clearing user data...")
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.createdAt = time.Time{}
	fmt.Println("User data cleared.")
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Time{},
	}, nil
}