package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const accountBalanceFile = "examples/bank/balances/balances.txt"

func validateAmount(amount float64) bool {
	if amount < 0 {
		return false
	}
	return true
}

func saveBalanceInFile(balance float64) {
	err := os.WriteFile(accountBalanceFile, []byte(fmt.Sprintf("%.2f", balance)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadBalanceFromFile() float64 {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		log.Fatal(err)
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		log.Fatal(err)
	}

	return balance
}

func goBank() {
	var choice int
	var balance, insertedAmount float64
	if os.ErrNotExist == nil {
		balance = LoadBalanceFromFile()
	}
	fmt.Println("Welcome to Go bank!")
	for {
		fmt.Println("---------GO BANK-----------")

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Println("---------GO BANK-----------")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("---------GO BANK-----------")
			fmt.Printf("Your balance: %.2f \n", balance)
		case 2:
			fmt.Println("---------GO BANK-----------")
			fmt.Print("Enter amount: ")
			fmt.Scan(&insertedAmount)
			if !validateAmount(insertedAmount) {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}
			balance += insertedAmount
			fmt.Printf("Balance updated! New balance: %.2f \n", balance)
		case 3:
			fmt.Println("---------GO BANK-----------")
			fmt.Print("Enter amount: ")
			fmt.Scan(&insertedAmount)
			if !validateAmount(insertedAmount) {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}

			if balance < insertedAmount {
				fmt.Println("You cannot withdraw money!")
				continue
			}
			balance -= insertedAmount

			fmt.Printf("Balance updated! New balance: %.2f \n", balance)
		case 4:
			fmt.Println("---------GO BANK-----------")
			fmt.Println("Bye :)")
			fmt.Println("---------------------------")
			saveBalanceInFile(balance)
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func main() {
	goBank()
}
