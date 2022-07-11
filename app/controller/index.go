package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/waline-async-mail/app"
	"log"
)

// Redirect
// GET /
func Redirect(c *gin.Context) {
	c.Redirect(302, "https://github.com/soxft/waline-async-mail")
}

// Handler
// POST /
func Handler(c *gin.Context) {
	var data app.CommentStruct
	err := c.ShouldBindJSON(&data)
	if err != nil {
		invalidRequest(c)
		return
	} else if data.Type != "new_comment" {
		invalidRequest(c)
		return
	}
	log.Println(data.Data)

	c.JSON(202, gin.H{
		"success": true,
		"message": "success",
		"data":    gin.H{},
	})
	return
}

func invalidRequest(c *gin.Context) {
	c.JSON(403, gin.H{
		"success": false,
		"message": "Invalid params",
		"data":    gin.H{},
	})
}
