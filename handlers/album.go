package handlers

import (
	"learn_api_test/db"
	"learn_api_test/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// func CreateAlbum(c echo.Context) error {}

func DeleteAlbum(c echo.Context) error {
	rqID, convertionErr := uuid.Parse(c.Param("ID"))
	if convertionErr != nil {
		err := models.NewError("convertion error", "not possible to converte UUID", http.StatusBadRequest, convertionErr)
		return c.JSON(http.StatusBadRequest, err)
	}

	var albunsSize int = len(db.Albums)
	for albumIndex, album := range db.Albums {
		if rqID == album.ID{
			db.Albums = append(db.Albums[:albumIndex], db.Albums[albumIndex+1:]...)			
		}
	}
	if albunsSize == len(db.Albums) {
		err := models.NewError("bad request", "not possible to find any album with this ID", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, err)
	}

	message := models.NewCommomMessage("album deletado", "album deletado com sucesso!")

	return c.JSON(http.StatusOK, message)
}

func GetAlbum(c echo.Context) error {
	rqID, convertionErr := uuid.Parse(c.Param("ID"))
	if convertionErr != nil {
		err := models.NewError("convertion error", "not possible to converte UUID", http.StatusBadRequest, convertionErr)
		return c.JSON(http.StatusBadRequest, err)
	}

	var album models.Album
	for _, albumDB := range db.Albums{
		if rqID == albumDB.ID {
			album = *models.NewAlbum(albumDB.ID, albumDB.Name)
		}
	}
	if album == (models.Album{}) {
		err := models.NewError("bad request", "not possible to find any album with this ID", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, album)
}

func GetAlbumsPerArtist(c echo.Context) error {
	rqID, convertionErr := uuid.Parse(c.Param("ID"))
	if convertionErr != nil {
		err := models.NewError("convertion error", "not possible to converte UUID", http.StatusBadRequest, convertionErr)
		return c.JSON(http.StatusBadRequest, err)
	}

	albumsIDs := make([]uuid.UUID, 0)
	artistFound := false
	for _, artist := range db.Artists {
		if artist.ID == rqID {
			artistFound = true
			albumsIDs = append(albumsIDs, artist.AlbumsIDs...)
		}
	}
	if !artistFound {
		err := models.NewError("bad request", "not possible to find any artist with this ID", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, err)
	}
	if len(albumsIDs) == 0 {
		err := models.NewError("not found", "not possible to find any album ID", http.StatusNotFound, nil)
		return c.JSON(http.StatusNotFound, err)
	}

	albums := make([]models.Album, 0)
	for _, album := range db.Albums {
		for _, albumID := range albumsIDs {
			if album.ID == albumID{
				foundAlbum := models.NewAlbum(album.ID, album.Name)
				albums = append(albums, *foundAlbum)
			}
		}
	}
	if len(albums) == 0 {
		err := models.NewError("convertion error", "not possible convert albumID to models.Album", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, albums)
}

// func UpdateAlbum(c echo.Context) error {}