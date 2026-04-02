package infrastructure

import (
	"encoding/json"
	"errors"
	"os"
	"problemSolving/examples/to-do/internal/note/domain"
	"problemSolving/examples/to-do/internal/note/dto"
)

type FileRepository struct {
	filePath string
}

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{
		filePath: path,
	}
}

func (r *FileRepository) Save(n *domain.Note) error {

	if n == nil {
		return errors.New("invalid note")
	}

	notes, err := r.load()
	if err != nil {
		return err
	}

	notes = append(notes, n)

	notesDTO := dto.ToDTOList(notes)

	data, err := json.MarshalIndent(notesDTO, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *FileRepository) load() ([]*domain.Note, error) {
	_, err := os.Stat(r.filePath)

	if errors.Is(err, os.ErrNotExist) {
		return []*domain.Note{}, nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []*domain.Note{}, nil
	}
	var notes []*domain.Note

	if err := json.Unmarshal(data, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}
