package db

import (
	"learn_api_test/models"

	"github.com/google/uuid"
)


var Albums = []models.Album{
	{
		ID: uuid.MustParse("108bca2e-77c7-4120-82bf-aae12884e3ec"),
		Name: "donda",
	},
	{
		ID: uuid.MustParse("11b9da2d-f9d7-4ecf-bb6a-10b0d8594e19"),
		Name: "the life of pablo",
	},
	{
		ID: uuid.MustParse("a32aa7f7-c338-4c01-918d-8e8e250e1529"),
		Name: "set me free",
	},
}