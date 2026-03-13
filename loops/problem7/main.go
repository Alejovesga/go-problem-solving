package main

import "fmt"

func main() {
	reverse()
}

func reverse() {
	var x, y int
	fmt.Scan(&x)
	for x > 0 {
		y *= 10
		y += x % 10
		x /= 10
	}
	fmt.Print(y)
}
