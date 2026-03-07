package main

import "fmt"

func main() {
	var choice int
	var balance, insertedAmount float64

	for {
		fmt.Println("---------GO BANK-----------")
		fmt.Println("Welcome to Go bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Println("---------GO BANK-----------")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Print("\033[H\033[2J")
		if choice == 1 {
			fmt.Println("---------GO BANK-----------")
			fmt.Printf("Your balance: %.2f \n", balance)
		} else if choice == 2 {
			fmt.Println("---------GO BANK-----------")
			fmt.Print("Enter amount: ")
			fmt.Scan(&insertedAmount)
			balance += insertedAmount
		} else if choice == 3 {
			fmt.Println("---------GO BANK-----------")
			fmt.Print("Enter amount: ")
			fmt.Scan(&insertedAmount)
			balance -= insertedAmount
		} else if choice == 4 {
			fmt.Println("---------GO BANK-----------")
			fmt.Println("Bye :)")
			fmt.Print("\033[H\033[2J")
			fmt.Println("---------------------------")
			break
		}
	}

}
