package main

import "fmt"

func main() {
	var x, firstDigit, secondDigit, thirdDigit int

	fmt.Print("Enter a number: ")
	fmt.Scan(&x)

	firstDigit = x / 100
	secondDigit = (x % 100) / 10
	thirdDigit = x % 10

	fmt.Println(firstDigit, secondDigit, thirdDigit)

}
