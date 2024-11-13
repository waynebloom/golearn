package main

import (
	"fmt"
  "golearn/structs/user"
)

func main() {
	firstName := promptAndScan("Please enter your first name: ")
	lastName := promptAndScan("Please enter your last name: ")
	birthDate := promptAndScan("Please enter your birthdate (MM/DD/YYYY): ")
  appUser := user.New(firstName, lastName, birthDate)
  appUser.Print()
}

func promptAndScan(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
