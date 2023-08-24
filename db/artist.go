package db

import (
	"learn_api_test/models"

	"github.com/google/uuid"
)

var artists = []models.Artist{
	{
		ID: uuid.MustParse("5da750ed-df00-46f1-9845-531a5b2ca597"),
		Name: "Kanye West",
		Members: []string{
			"Kanye West Junior",
		},
		AlbumsIDs: []uuid.UUID{
			uuid.MustParse("108bca2e-77c7-4120-82bf-aae12884e3ec"),
			uuid.MustParse("11b9da2d-f9d7-4ecf-bb6a-10b0d8594e19"),
		},
	},
	{
		ID: uuid.MustParse("ff2cf81e-2e46-4510-bb8e-ec1f663c319d"),
		Name: "Twice",
		Members: []string{
			"Nayeon",
			"Jeongyeon",
			"Momo",
			"Sana",
			"Jihyo",
			"Mina",
			"Dahyun",
			"Chaeyoung",
			"Tzuyu",
		},
		AlbumsIDs: []uuid.UUID{
			uuid.MustParse("a32aa7f7-c338-4c01-918d-8e8e250e1529"),
		},
	},
}