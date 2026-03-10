package main

import "fmt"

func main() {
	determineEvenOdd()
}

func determineEvenOdd() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
