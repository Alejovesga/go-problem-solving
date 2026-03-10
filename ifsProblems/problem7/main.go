package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := calculator()
	if err != nil {
		panic(err)
	} else {
		fmt.Print(result)
	}
}

func calculator() (float64, error) {
	var a, b float64
	var op string
	fmt.Scan(&a, &op, &b)

	if op == "+" {
		return a + b, nil
	} else if op == "-" {
		return a - b, nil
	} else if op == "*" {
		return a * b, nil
	} else if op == "/" {
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return a / b, nil
	}
	return 0, errors.New("invalid operation")

}
