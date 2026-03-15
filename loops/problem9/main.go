package main

import "fmt"

func main() {
	colatz()
}

func colatz() {
	var x, count int
	fmt.Scan(&x)
	for x > 1 {
		if x%2 == 0 {
			x = x / 2
			count++
			continue
		}
		x = 3*x + 1
		count++
	}
	fmt.Println(count)
}
