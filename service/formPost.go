package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// content-type:application/x-www-form-urlencoded
func FormPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
