package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Enter x: ")
	fmt.Scan(&x)
	fmt.Print("Enter y: ")
	fmt.Scan(&y)

	x, y = y, x

	fmt.Printf("x: %d, y: %d", x, y)
}
