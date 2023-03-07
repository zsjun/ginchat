package router

import (
	"ginchat/common"
	"ginchat/docs"
	"ginchat/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AuthRequired is a simple middleware to check the session.
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(common.Userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func Router() *gin.Engine {
	r := gin.Default()
	// Set up the session store
	// Setup the cookie store for session management
	r.Use(sessions.Sessions("mysession", cookie.NewStore(common.Secret)))

	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/create", service.CreateUser)
	r.PUT("/user/update", service.UpdateUser)
	r.DELETE("/user/delete", service.DeleteUser)

	// Private group, require authentication to access
	// private := r.Group("/private")
	// private.Use(AuthRequired)
	// {
	// 	private.GET("/me", me)
	// 	private.GET("/status", status)
	// }
	return r

}
