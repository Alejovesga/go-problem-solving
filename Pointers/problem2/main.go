package main

import "fmt"

func main() {
	var option string
	var count, amount int

	fmt.Scanln(&option, &amount)

	for option != "" {
		if option == "increment" {
			increment(&count, amount)
		} else {
			decrement(&count, amount)
		}
		fmt.Println(count)
		fmt.Scanln(&option, &amount)
	}

}

func increment(count *int, amount int) {
	*count += amount
}

func decrement(count *int, amount int) {
	*count -= amount
}
