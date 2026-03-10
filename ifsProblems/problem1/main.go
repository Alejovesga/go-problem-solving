package main

import "fmt"

func main() {
	determinePairOdd()
}

func determinePairOdd() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
}
