package webutil

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/waline-async-mail/app/controller"
)

func initRoute(r *gin.Engine) {
	r.GET("/ping", controller.Ping)
	r.GET("/", controller.Index)
	r.POST("/", controller.Handler)
}
