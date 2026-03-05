package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159
	var radius, perimeter, area float64
	fmt.Print("Ingresa el radio del circulo")
	fmt.Scan(&radius)

	perimeter = 2 * pi * radius
	area = pi * math.Pow(radius, 2)

	fmt.Println("El perimetro:", perimeter)
	fmt.Println("El area:", area)

}
