package main

import (
	"fmt"
	"problemSolving/examples/to-do/internal/infrastructure"
	"problemSolving/examples/to-do/internal/note/application"
)

const path string = "./storage/notes"

func main() {
	fmt.Println("--To do App--")
	repo := infrastructure.NewFileRepository(path)
	createNote := application.NewCreateNote(repo)

	err := createNote.Execute("test", "content test")

	if err != nil {
		fmt.Println(err)
	}
}
