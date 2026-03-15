package main

import "fmt"

func main() {
	triangleAsterik()
}

func triangleAsterik() {
	var x int
	var str string
	fmt.Scan(&x)
	for x > 0 {
		str = str + "*"
		fmt.Println(str)
		x -= 1
	}
}
