package main

import (
	"learn_api_test/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	albumRoutes := e.Group("/album")
	// artistRoutes := e.Group("/artist")
	// songRoutes := e.Group("/song")

	albumRoutes.GET("/get/all/artist/:ID", handlers.GetAlbumsPerArtist)
	albumRoutes.GET("/get/:ID", handlers.GetAlbum)
	// albumRoutes.POST("/create", handlers.CreateAlbum)
	// albumRoutes.PUT("/update/:ID", handlers.UpdateAlbum)
	albumRoutes.DELETE("/delete/:ID", handlers.DeleteAlbum)

	// artistRoutes.GET("/get/all", handlers.GetArtists)
	// artistRoutes.GET("/get/:ID", handlers.GetArtist)
	// artistRoutes.POST("/create", handlers.CreateArtist)
	// artistRoutes.PUT("/update/:ID", handlers.UpdateArtist)
	// artistRoutes.DELETE("/delete/:ID", handlers.DeleteArtist)

	// songRoutes.GET("/get/all/artist/:ID", handlers.GetSongs)
	// songRoutes.GET("/get/:ID", handlers.GetSong)
	// songRoutes.POST("/create", handlers.CreateSong)
	// songRoutes.PUT("/update/:ID", handlers.UpdateSong)
	// songRoutes.DELETE("/delete/:ID", handlers.DeleteSong)

	e.Logger.Fatal((e.Start(":8000")))
}
