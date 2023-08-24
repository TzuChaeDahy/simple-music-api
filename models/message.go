package models

type CommomMessage struct {
	Name string
	Description string
}

func NewCommomMessage(name string, description string) *CommomMessage {
	return &CommomMessage{
		Name: name,
		Description: description,
	}
}