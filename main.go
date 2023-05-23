package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitq50/go-lang-project/controller"
	"github.com/rohitq50/go-lang-project/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func main() {
	// Initialize config
	// conf := config.NewConfig()

	server := gin.Default()

	server.GET("/users", func(ctxhttp *gin.Context) {
		ctxhttp.JSON(200, userController.FindAll())
	})

	server.POST("/users", func(ctxhttp *gin.Context) {
		ctxhttp.JSON(200, userController.Save(ctxhttp))
	})

	server.Run(":8080")
}
