package main

import "fmt"

func main() {
	fmt.Print(rockPaperScissors())
}

func rockPaperScissors() string {
	var player1, player2 string
	fmt.Scan(&player1, &player2)

	if player1 == "rock" && player2 == "scissors" {
		return "player1"
	} else if player1 == "scissors" && player2 == "rock" {
		return "player2"
	} else if player1 == "scissors" && player2 == "paper" {
		return "player1"
	} else if player1 == "paper" && player2 == "scissors" {
		return "player2"
	} else if player1 == "paper" && player2 == "rock" {
		return "player1"
	} else if player1 == "rock" && player2 == "paper" {
		return "player2"
	}
	return "draw"
}
