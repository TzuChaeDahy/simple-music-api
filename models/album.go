package models

type Album struct {
	Name string
}

func NewAlbum(name string) *Album {
	return &Album{
		Name: name,
	}
}