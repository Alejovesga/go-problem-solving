package main

import "fmt"

func main() {
	allEven()
}

func allEven() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
