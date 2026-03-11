package main

import "fmt"

func main() {
	primal()
}

func primal() {
	var x, count int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
