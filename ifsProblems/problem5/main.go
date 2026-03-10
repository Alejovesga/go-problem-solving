package main

import "fmt"

func main() {
	fmt.Print(traffic())
}

func traffic() string {
	var color string
	fmt.Println("Enter a color:")
	fmt.Scan(&color)

	if color == "red" {
		return "stop"
	} else if color == "yellow" {
		return "precaution"
	} else if color == "green" {
		return "go"
	}
	return "not valid color"

}
