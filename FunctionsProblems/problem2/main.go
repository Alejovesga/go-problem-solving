package main

import (
	"errors"
	"fmt"
	"os"
)

const dataFilePath = "FunctionsProblems/problem2/storage/data.txt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeCalculatorInformation(expenses, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		panic(err)
	} else if userInput == 0 {
		err := errors.New("user input error: the value cannot be zero")
		panic(err)
	} else if userInput < 0 {
		err := errors.New("user input error: the value cannot be negative")
		panic(err)
	}
	return userInput
}

func storeCalculatorInformation(revenue, expenses, taxRate float64) {
	data := fmt.Sprintf("revenue : %.2f \nexpenses: %.2f \ntaxRate: %.2f \n", revenue, expenses, taxRate)
	err := os.WriteFile(dataFilePath, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
