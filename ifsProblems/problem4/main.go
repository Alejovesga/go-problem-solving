package main

import "fmt"

func main() {
	fmt.Print(isLeap())
}

func isLeap() bool {
	var year int
	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
