package models

type Artist struct {
	Name string
	Members []string
	Albums []Album
}

func NewArtist(name string, members []string, albums []Album) *Artist {
	return &Artist{
		Name: name,
		Members: members,
		Albums: albums,
	}
}