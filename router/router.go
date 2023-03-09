package router

import (
	"ginchat/common"
	"ginchat/docs"
	"ginchat/service"
	"net/http"

	"github.com/gin-contrib/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AuthRequired is a simple middleware to check the session.
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(common.Userkey)
	if userId == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not logined"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func Router() *gin.Engine {
	// default
	// r := gin.Default()
	// Creates a router without any middleware by default
	// use middware
	r := gin.New()
	// 设置生成sessionId的密钥
	store := cookie.NewStore([]byte("secret"))
	// mysession是返回給前端的sessionId名
	r.Use(sessions.Sessions("mysession", store))

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	authMiddleware := auth.BasicAuth(AuthRequired)
	// Initialize basic auth middleware
	// docs Swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	// Simple group: v1
	r.POST("/login", service.Login)

	v1 := r.Group("/api", authMiddleware)
	{
		v1.GET("/getUserList", service.GetUserList)
		v1.POST("/create", service.CreateUser)
		v1.PUT("/update", service.UpdateUser)
		v1.DELETE("/delete", service.DeleteUser)
		v1.GET("/:name", service.GetRouterName)
		r.GET("/logout", service.LogOut)
		// Set a lower memory limit for multipart forms (default is 32 MiB)
		r.MaxMultipartMemory = 50 << 20 // 50 MiB
		r.POST("/upload", service.Upload)
		r.POST("/testJson", service.TestPostMethod)
		r.GET("/welcome", service.Welcome)
		r.POST("/form_post", service.FormPost)
		r.GET("/isLogin", service.IsLogin)
		r.GET("/ping", service.Ping)
	}

	return r

}
