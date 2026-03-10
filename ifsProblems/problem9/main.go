package main

import "fmt"

func main() {
	fmt.Print(calificationConverter())
}

func calificationConverter() string {
	var calification int
	fmt.Scan(&calification)

	if calification > 100 || calification < 0 {
		return "Invalid"
	} else if calification >= 90 {
		return "A"
	} else if calification >= 80 {
		return "B"
	} else if calification >= 70 {
		return "C"
	} else if calification >= 60 {
		return "D"
	}
	return "F"
}
