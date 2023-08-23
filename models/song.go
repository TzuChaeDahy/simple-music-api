package models

import "time"

type Song struct {
	Name string
	Artist Artist
	Feat []Artist
	Duration time.Duration
	Album Album
}

func NewSong(name string, artist Artist, feat []Artist, duration time.Duration, album Album) *Song {
	return &Song{
		Name: name,
		Artist: artist,
		Feat: feat,
		Duration: duration,
		Album: album,
	}
}