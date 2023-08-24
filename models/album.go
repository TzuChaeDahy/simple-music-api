package models

import "github.com/google/uuid"

type Album struct {
	ID uuid.UUID
	Name string
}

func NewAlbum(id uuid.UUID, name string) *Album {
	return &Album{
		ID: id,
		Name: name,
	}
}