package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitq50/go-lang-project/entity"
	"github.com/rohitq50/go-lang-project/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctxhttp *gin.Context) entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}
func (c *controller) Save(ctxhttp *gin.Context) entity.User {
	var user entity.User
	ctxhttp.BindJSON(&user)
	c.service.Save(user)
	return user
}
