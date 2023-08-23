package models

type Artist struct {
	ID int
	Name string
	Members []string
	AlbumsIDs []int
}

func NewArtist(id int, name string, members []string, albums []int) *Artist {
	return &Artist{
		ID: id,
		Name: name,
		Members: members,
		AlbumsIDs: albums,
	}
}