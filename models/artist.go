package models

import "github.com/google/uuid"

type Artist struct {
	ID uuid.UUID
	Name string
	Members []string
	AlbumsIDs []uuid.UUID
}

func NewArtist(id uuid.UUID, name string, members []string, albums []uuid.UUID) *Artist {
	return &Artist{
		ID: id,
		Name: name,
		Members: members,
		AlbumsIDs: albums,
	}
}