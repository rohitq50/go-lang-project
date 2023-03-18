package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitq50/go-lang-project/entity"
	"github.com/rohitq50/go-lang-project/service"
)

type VideoController interface {
	FindAll() []entity.Videos
	Save(ctxhttp *gin.Context) entity.Videos
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Videos {
	return c.service.FindAll()
}
func (c *controller) Save(ctxhttp *gin.Context) entity.Videos {
	var video entity.Videos
	ctxhttp.BindJSON(&video)
	c.service.Save(video)
	return video
}
