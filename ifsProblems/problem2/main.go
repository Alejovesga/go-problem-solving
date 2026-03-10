package main

import "fmt"

func main() {
	greater := greater()
	fmt.Println("Greater is ", greater)
}

func greater() int {
	var a, b, c int
	fmt.Print("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	if a > b && b > c {
		return a
	} else if b > c {
		return b
	}
	return c
}
