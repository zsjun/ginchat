package router

import (
	"fmt"
	"ginchat/docs"
	"ginchat/global"
	"ginchat/models"
	"ginchat/service"
	"ginchat/ws"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AuthRequired is a simple middleware to check the session.
func authMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(global.Userkey)
	if userId == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you are not logined"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
func tokenAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.UserBasic{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return global.JwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*models.UserBasic)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("username", claims.UserName)
	c.Next()
}

func Router() *gin.Engine {
	// default
	// r := gin.Default()
	// Creates a router without any middleware by default
	// use middware
	r := gin.New()
	// secure
	r.SecureJsonPrefix(")]} ',\n")
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
	// Initialize basic auth middleware
	// docs Swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Simple group: v1
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 50 << 20 // 50 MiB
	r.POST("/login", service.TokenLogin)
	r.GET("/api/ws", ws.WsHandle)
	v1 := r.Group("/api", tokenAuthMiddleware)
	{
		v1.GET("/getUserList", service.GetUserList)
		v1.POST("/create", service.CreateUser)
		v1.PUT("/update", service.UpdateUser)
		v1.DELETE("/delete", service.DeleteUser)
		v1.GET("/:name", service.GetRouterName)
		v1.POST("/upload", service.Upload)
		v1.POST("/testJson", service.TestPostMethod)
		v1.GET("/welcome", service.Welcome)
		v1.POST("/form_post", service.FormPost)

		// v1.GET("/isLogin", service.IsLogin)
	}

	r.GET("/logout", service.TokenLogOut)

	return r

}
