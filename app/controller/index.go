package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/waline-async-mail/app"
	"github.com/soxft/waline-async-mail/config"
	"github.com/soxft/waline-async-mail/library/mail"
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

	if err != nil || data.Type != "new_comment" {
		c.JSON(403, gin.H{
			"success": false,
			"message": "Invalid params",
			"data":    gin.H{},
		})
		return
	}

	if data.Data.Reply.Status == "" {
		// 评论邮件 > 发送给 Owner
		var ownerArgs mail.OwnerArgs
		ownerArgs = mail.OwnerArgs{
			Author:    data.Data.Comment.Nick,
			Permalink: config.BlogInfo.Addr + data.Data.Comment.Url,
			SiteTitle: config.BlogInfo.Title,
			Ip:        data.Data.Comment.Ip,
			Time:      data.Data.Comment.InsertedAt,
			Status:    data.Data.Comment.Status,
			Mail:      data.Data.Comment.Mail,
		}
		content, err := mail.ParseOwner(ownerArgs)
		log.Println(content, err)
	} else {
		// 回复邮件 > 发送给 Owner & 被回复者
	}

	c.JSON(202, gin.H{
		"success": true,
		"message": "success",
		"data":    gin.H{},
	})
	return
}
