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
	r.POST("/login", service.Login)
	r.GET("/ping", service.Ping)
	r.GET("/user/:name", service.GetRouterName)

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to an url matching:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// content-type:application/x-www-form-urlencoded
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	return r

}
