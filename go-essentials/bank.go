package main

import (
	"fmt"
	"strconv"
  "golearn/bank/fileops"
  "github.com/Pallinder/go-randomdata"
)

const withdrawalFee = 2.99
const balanceFile = "balance.txt"

func main() {
	var balance, err = fileops.ReadFloat(balanceFile)

  if err != nil {
    fmt.Println("Encountered an error while reading data from file:")
    fmt.Println(err)
    fmt.Println("--------------")
  }

	fmt.Println("Welcome to Go Bank")
  fmt.Println(randomdata.PhoneNumber())

	for {
		// visual option menu
		fmt.Print("What do you want to do?\n1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit\n")

		// prompt for input
		var choice string
		fmt.Print("Select an option: ")
		fmt.Scan(&choice)
		fmt.Print("\n---\n\n")

		switch choice {
		case "1":
			fmt.Printf("Your balance is %.2f", balance)
		case "2":
			balance += deposit()
			fmt.Printf("Your balance is now %.2f\n", balance)
		case "3":
			balance -= (withdraw(balance) + withdrawalFee)
			fmt.Printf("Your balance is now %.2f\n", balance)
		case "4":
			fmt.Println("Thank you for using Go Bank. Goodbye!")
      fileops.WriteFloat(balanceFile, balance)
			return
		default:
			fmt.Printf("ERROR: %v is not a valid selection.\n", choice)
      break
		}

		fmt.Print("\n---\n\n")
	}
}

func deposit() (value float64) {
	var input string
	fmt.Println("How much would you like to deposit?")

	for value == 0 {
		fmt.Println("Desired deposit: ")
		fmt.Scan(&input)
		floatNum, error := strconv.ParseFloat(input, 64)

		if error != nil {
			fmt.Println("ERROR: Please enter a valid currency value.")
		} else if floatNum <= 0.0 {
			fmt.Println("ERROR: Please enter a positive, non-zero value.")
		} else {
			value = floatNum
		}
	}

	fmt.Printf("You successfully deposited %.2f.\n", value)
	return
}

func withdraw(maxWithdrawal float64) (value float64) {
	var input string
	fmt.Printf("All withdrawals from Go Bank are subject to the %v Blatantly Excessive Withdrawal Fee\n", withdrawalFee)
	fmt.Println("How much would you like to withdraw?")

	for value == 0 {
		fmt.Println("Desired amount: ")
		fmt.Scan(&input)
		floatNum, error := strconv.ParseFloat(input, 64)

		if error != nil {
			fmt.Println("ERROR: Please enter a valid amount.")
		} else if floatNum <= 0.0 {
			fmt.Println("ERROR: Please enter a positive, non-zero amount.")
		} else if floatNum > maxWithdrawal-withdrawalFee {
			fmt.Println("ERROR: You cannot withdraw more than you have. That would be stealing.")
		} else {
			value = floatNum
		}
	}

	fmt.Printf("You successfully withdrew %.2f.\n", value)
	return
}
