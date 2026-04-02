package domain

import (
	"errors"
	"time"
)

type Note struct {
	title   string
	content string
	date    time.Time
}

func NewNote(title, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if content == "" {
		return nil, errors.New("content cannot be empty")
	}
	return &Note{
		title:   title,
		content: content,
		date:    time.Now(),
	}, nil
}

func (n *Note) Title() string {
	return n.title
}

func (n *Note) Content() string {
	return n.content
}

func (n *Note) Date() time.Time {
	return n.date
}
