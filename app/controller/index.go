package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type PostStruct struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// Index
// GET /
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Handler
// POST /
func Handler(c *gin.Context) {
	var data PostStruct
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		c.JSON(403, gin.H{
			"success": false,
			"message": "Invalid request",
			"data":    gin.H{},
		})
		return
	}
	log.Println(data)
	comment := data.Data["comment"]
	log.Println(comment)
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
