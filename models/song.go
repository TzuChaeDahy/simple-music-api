package models

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID           uuid.UUID
	Name         string
	ArtistID     uuid.UUID
	FeatArtistID []uuid.UUID
	Duration     time.Duration
	AlbumID      uuid.UUID
}

func NewSong(id uuid.UUID, name string, artistID uuid.UUID, featArtistID []uuid.UUID, duration time.Duration, albumID uuid.UUID) *Song {
	return &Song{
		ID:           id,
		Name:         name,
		ArtistID:     artistID,
		FeatArtistID: featArtistID,
		Duration:     duration,
		AlbumID:      albumID,
	}
}
