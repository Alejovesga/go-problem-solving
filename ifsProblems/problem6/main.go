package main

import "fmt"

func main() {
	fmt.Print(clasificateTriangle())
}

func clasificateTriangle() string {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (a+b) > c && (a+c) > b && (b+c) > a {
		if a == b && b == c {
			return "Equilateral"
		} else if (a == b) || (b == c) || (c == a) {
			return "Isosceles"
		}
		return "escaleno"
	} else {
		return "Not triangle"
	}
}
