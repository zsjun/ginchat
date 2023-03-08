package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// uppercase export
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
