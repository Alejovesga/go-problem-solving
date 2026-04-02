package application

import (
	"problemSolving/examples/to-do/internal/note/domain"
)

type CreateNote struct {
	repo domain.Repository
}

func NewCreateNote(repo domain.Repository) *CreateNote {
	return &CreateNote{
		repo: repo,
	}
}

func (c *CreateNote) Execute(title, content string) error {
	note, err := domain.NewNote(title, content)
	if err != nil {
		return err
	}

	return c.repo.Save(note)
}
