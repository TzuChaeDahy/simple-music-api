package models

import "time"

type Song struct {
	ID int
	Name string
	ArtistID int
	FeatArtistID []int
	Duration time.Duration
	AlbumID int
}

func NewSong(id int, name string, artistID int, featArtistID []int, duration time.Duration, albumID int) *Song {
	return &Song{
		ID: id,
		Name: name,
		ArtistID: artistID,
		FeatArtistID: featArtistID,
		Duration: duration,
		AlbumID: albumID,
	}
}