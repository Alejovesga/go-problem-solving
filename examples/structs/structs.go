package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	var appUser User

	appUser = User{
		createdAt: time.Now(),
	}

	appUser.firstName = getUserData("Please enter your first name: ")
	appUser.lastName = getUserData("Please enter your last name: ")
	appUser.birthDate = getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	outputUserDetails(&appUser)
}

func outputUserDetails(user *User) {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
