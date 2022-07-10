package webutil

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	r := gin.Default()

	initRoute(r)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
