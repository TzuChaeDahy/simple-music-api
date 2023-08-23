package models

type Album struct {
	ID int
	Name string
}

func NewAlbum(id int, name string) *Album {
	return &Album{
		ID: id,
		Name: name,
	}
}