package main

import "fmt"

func main() {
	result := sum()
	fmt.Println(result)
}

func sum() int {
	var x, result int
	fmt.Scan(&x)
	for i := 0; i <= x; i++ {
		result += i
	}
	return result
}
