package main

import "fmt"

func main() {
	regresive()
}

func regresive() {
	var x int
	fmt.Scan(&x)
	for x > 0 {
		fmt.Println(x)
		x--
	}
	fmt.Println("Go!")
}
