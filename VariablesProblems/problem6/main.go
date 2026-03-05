package main

import "fmt"

func main() {

	var tip, bill, diners, totalBill, payPerson float64

	fmt.Print("Enter the tip: ")
	fmt.Scan(&tip)
	fmt.Print("Enter the bill: ")
	fmt.Scan(&bill)
	fmt.Print("Enter the diners: ")
	fmt.Scan(&diners)

	totalBill = bill + (bill * (tip / 100))
	payPerson = totalBill / diners

	fmt.Printf("pay by person: %f", payPerson)

}
