package main

import "fmt"

func main() {
	factorial()
}

func factorial() {
	var x int
	result := 1
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		result *= i
	}
	fmt.Println(result)
}
