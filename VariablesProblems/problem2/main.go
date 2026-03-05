package main

import "fmt"

func main() {
	var celsius, fahrenheit, kelvin float64

	fmt.Print("Enter your celsius: ")
	fmt.Scan(&celsius)

	fahrenheit = (celsius * 9 / 5) + 32
	kelvin = celsius + 279.15

	fmt.Printf("Celsius: %f, Fahrenheit: %f, kelvin: %f", celsius, fahrenheit, kelvin)
}
