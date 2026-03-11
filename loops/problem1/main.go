package main

import "fmt"

func main() {
	calculator()
}

func calculator() {
	var x int
	fmt.Scan(&x)
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", x, i, x*i)
	}
}
