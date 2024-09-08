package types

type Note struct {
	Id          string `json:id`
	Name        string `json:name`
	Description string `json:description`
}

func NewNote(id, name, description string) *Note {
	return &Note{
		Id:          id,
		Name:        name,
		Description: description,
	}
}
