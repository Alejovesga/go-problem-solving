package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(passwordValidator())
}

func passwordValidator() string {
	var password, notValid string
	fmt.Scan(&password, &notValid)

	if len(password) >= 8 {
		if strings.ContainsAny(password, "0123456789") {
			if len(notValid) > 0 {
				return "invalid"
			}
			return "valid"
		}
		return "invalid"
	}
	return "invalid"
}
