package webutil

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/waline-async-mail/config"
	"log"
)

func Init() {
	if config.Server.Debug {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	initRoute(r)

	log.Printf("[INFO] server is running at %s", config.Server.Addr)
	err := r.Run(config.Server.Addr)
	if err != nil {
		log.Fatal(err)
	}
}
