package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Query string parameters are parsed using the existing underlying request object.
// The request responds to an url matching:  /welcome?firstname=Jane&lastname=Doe
func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	// shortcut for c.Request.URL.Query().Get("lastname")
	lastname := c.Query("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
