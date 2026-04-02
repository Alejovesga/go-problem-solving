package dto

import (
	"problemSolving/examples/to-do/internal/note/domain"
	"time"
)

type NoteDTO struct {
	Title   string
	Content string
	Date    time.Time
}

func toDTO(n *domain.Note) NoteDTO {
	return NoteDTO{
		Title:   n.Title(),
		Content: n.Content(),
		Date:    n.Date(),
	}
}

func ToDTOList(notes []*domain.Note) []NoteDTO {
	dtos := make([]NoteDTO, len(notes))

	for i, n := range notes {
		dtos[i] = toDTO(n)
	}

	return dtos
}
