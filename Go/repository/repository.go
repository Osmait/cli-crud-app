package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/osmait/cli-crud-app/types"
)

type NoteRepository struct {
	FilePath string
	NamePath string
}

func NewNoteRepository() (*NoteRepository, error) {
	filename := "./db.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		initialData := []types.Note{}
		initialDataJSON, err := json.Marshal(initialData)
		if err != nil {
			return nil, fmt.Errorf("error al codificar datos iniciales: %w", err)
		}

		file, err := os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("error al crear el archivo: %w", err)
		}
		defer file.Close()

		if _, err := file.Write(initialDataJSON); err != nil {
			return nil, fmt.Errorf("error al escribir datos iniciales en el archivo: %w", err)
		}
	}

	return &NoteRepository{}, nil
}

func (n *NoteRepository) Save(note types.Note) error {
	result, err := n.FindAll()
	if err != nil {
		return err
	}

	result = append(result, note)

	modif, err := json.MarshalIndent(result, "", "  ") // MarshalIndent para una mejor legibilidad
	if err != nil {
	}

	if err := os.WriteFile("./db.json", modif, 0644); err != nil {
		return err
	}

	return nil
}

func (n *NoteRepository) FindAll() ([]types.Note, error) {
	file, err := os.ReadFile("./db.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []types.Note{}, nil
		}
		return nil, err
	}

	var notes []types.Note
	if err := json.Unmarshal(file, &notes); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return notes, nil
}
