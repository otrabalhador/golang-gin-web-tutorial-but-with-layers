package web

import (
	"example/web-service-gin/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) addAlbumRoutes(router *gin.Engine) {
	router.GET("/albums/:id", app.getAlbum)
	router.GET("/albums", app.listAlbums)
	router.POST("/albums", app.createAlbum)
}

func (app *App) getAlbum(c *gin.Context) {
	id := c.Param("id")

	album := app.AlbumRepository.Get(id)
	if album != nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (app *App) listAlbums(c *gin.Context) {
	albums := app.AlbumRepository.List()

	c.IndentedJSON(http.StatusOK, albums)
}

func (app *App) createAlbum(c *gin.Context) {
	var newAlbum domain.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums := app.AlbumRepository.Create(newAlbum)

	c.IndentedJSON(http.StatusOK, albums)
}
