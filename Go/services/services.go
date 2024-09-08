package services

import (
	"github.com/osmait/cli-crud-app/repository"
	"github.com/osmait/cli-crud-app/types"
)

type NoteServices struct {
	noteRepository *repository.NoteRepository
}

func NewNoteServices(noteRepository *repository.NoteRepository) *NoteServices {
	return &NoteServices{
		noteRepository: noteRepository,
	}
}

func (n *NoteServices) Create(name, description string) {
	noteToSave := types.NewNote("1", name, description)
	n.noteRepository.Save(*noteToSave)
}

func (n *NoteServices) GetAll() {
	n.noteRepository.FindAll()
}
