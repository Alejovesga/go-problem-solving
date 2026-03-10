package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print(passwordValidator())
}

func passwordValidator() string {
	var password string

	reader := bufio.NewReader(os.Stdin)
	password, _ = reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	if len(password) >= 8 {
		if strings.Contains(password, " ") {
			return "invalid"
		}
		if strings.ContainsAny(password, "0123456789") {
			return "valid"
		}
		return "invalid"
	}
	return "invalid"
}
