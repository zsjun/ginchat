package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouterName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
