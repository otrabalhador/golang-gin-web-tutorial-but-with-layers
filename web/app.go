package web

import (
	"example/web-service-gin/repository"
	"github.com/gin-gonic/gin"
)

type App struct {
	AlbumRepository *repository.AlbumRepository
}

func NewApp(albumRepository *repository.AlbumRepository) *App {
	app := new(App)
	app.AlbumRepository = albumRepository
	return app
}

func (app *App) AddRoutes(router *gin.Engine) {
	app.addAlbumRoutes(router)
}
