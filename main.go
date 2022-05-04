package main

import (
	"example/web-service-gin/repository"
	"example/web-service-gin/web"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	albumRepository := repository.NewAlbumRepository()
	web.NewApp(albumRepository).
		AddRoutes(engine)

	engine.Run("localhost:8080")
}
