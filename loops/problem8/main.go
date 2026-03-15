package main

import "fmt"

func main() {
	fibonacci()
}

func fibonacci() {
	var x, fibo int
	z := 1
	fmt.Scan(&x)
	for fibo <= x {
		fmt.Println(fibo)
		fibo, z = z, z+fibo
	}
}
