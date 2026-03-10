package main

import "fmt"

func main() {

	fmt.Println(calculateImc())
}

func calculateImc() string {
	var weight, height float64

	fmt.Print("Enter weight and height: ")
	fmt.Scan(&weight, &height)

	imc := weight / (height * height)

	if imc < 18.5 {
		return "Low weight"
	} else if imc <= 25 {
		return "Normal weight"
	} else if imc < 30 {
		return "over weight"
	}
	return "Obesity"
}
