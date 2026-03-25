package main

import "fmt"

func main() {
	var a, b int
	reader(&a, &b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func reader(a, b *int) {
	fmt.Scan(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
