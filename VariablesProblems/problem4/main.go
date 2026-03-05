package main

import "fmt"

func main() {
	var seconds, hours, minutes int

	fmt.Print("Enter seconds: ")
	fmt.Scan(&seconds)

	hours = seconds / 3600
	minutes = (seconds % 3600) / 60
	seconds = (seconds % 3600) % 60
	fmt.Printf("Hours: %d Minutes: %d seconds: %d", hours, minutes, seconds)
}
