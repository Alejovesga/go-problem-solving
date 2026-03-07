package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	scanUserInteraction("Revenue: ", &revenue)
	scanUserInteraction("Expenses: ", &expenses)
	scanUserInteraction("Tax Rate: ", &taxRate)

	ebt, profit, ratio := calculateFutureValues(revenue, expenses, taxRate)

	printResults(ebt, profit, ratio)
}

func scanUserInteraction(message string, variable *float64) {
	fmt.Print(message)
	fmt.Scan(variable)
}

func calculateFutureValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func printResults(ebt, profit, ratio float64) {
	fmt.Printf("ebt: %.2f \nprofit: %.2f \nratio: %.2f", ebt, profit, ratio)
}
