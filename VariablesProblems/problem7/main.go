package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	outputText("Enter number of revenue: ")
	fmt.Scan(&revenue)
	outputText("Enter number of expenses: ")
	fmt.Scan(&expenses)
	outputText("Enter total tax rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFutureValues(revenue, expenses, taxRate)

	formatedText := fmt.Sprintf("EBT: %.2f \nProfit: %.2f \nRatio: %.2f \n", ebt, profit, ratio)

	fmt.Print(formatedText)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
