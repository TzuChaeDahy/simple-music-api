package db

import (
	"learn_api_test/models"
	"time"
	"github.com/google/uuid"
)

var songs = []models.Song{
	{
		ID: uuid.MustParse("9d9b380f-c11e-498a-8e28-f03d59ec88ec"),
		Name: "Saint Pablo",
		ArtistID: uuid.MustParse("5da750ed-df00-46f1-9845-531a5b2ca597"),
		FeatArtistID: []uuid.UUID{},
		Duration: time.Minute * 6 + time.Second * 45,
		AlbumID: uuid.MustParse("11b9da2d-f9d7-4ecf-bb6a-10b0d8594e19"),		
	},
	{
		ID: uuid.MustParse("12074c43-12d9-4c79-a735-d1e77d73498f"),
		Name: "Goals",
		ArtistID: uuid.MustParse("5da750ed-df00-46f1-9845-531a5b2ca597"),
		FeatArtistID: []uuid.UUID{},
		Duration: time.Minute * 3 + time.Second * 20,
		AlbumID: uuid.MustParse("108bca2e-77c7-4120-82bf-aae12884e3ec"),		
	},
	{
		ID: uuid.MustParse("6d657875-24a5-4801-829e-790476ef0cb6"),
		Name: "Set Me Free",
		ArtistID: uuid.MustParse("3866b178-1494-47d9-92b3-e0cecc4be172"),
		FeatArtistID: []uuid.UUID{},
		Duration: time.Minute * 3 + time.Second * 20,
		AlbumID: uuid.MustParse("a32aa7f7-c338-4c01-918d-8e8e250e1529"),		
	},
}