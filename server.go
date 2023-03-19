package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitq50/go-lang-project/controller"
	"github.com/rohitq50/go-lang-project/service"
)

var (
	videoService    service.UserService       = service.New()
	videoController controller.UserController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/users", func(ctxhttp *gin.Context) {
		ctxhttp.JSON(200, videoController.FindAll())
	})

	server.POST("/users", func(ctxhttp *gin.Context) {
		ctxhttp.JSON(200, videoController.Save(ctxhttp))
	})

	server.Run(":8080")
}
