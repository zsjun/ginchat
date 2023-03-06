package main

import (
	"fmt"
	"ginchat/router"
	"ginchat/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
  utils.InitConfig()
  utils.InitMysql()
  r := router.Router()
  fmt.Println("hello, world1")

  r.GET("/login", Login)
	
  // 需要登陆保护的
  auth := r.Group("/api")
	auth.Use(AuthRequired())
	{
		auth.GET("/me", UserInfo)
		auth.GET("/logout", Logout)
	}
  r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 登陆
func Login(c *gin.Context) {
  // 为了演示方便，我直接通过url明文传递账号密码，实际生产中应该用HTTP POST在body中传递
  userID := c.Query("user_id")
  password := c.Query("password")
  fmt.Println(userID,password)
  // 用户身份校验（查询数据库）
  if userID == "007" && password == "007" {
    // 生成cookie
    expiration := time.Now()
    expiration = expiration.AddDate(0, 0, 1)
    // 实际生产中我们可以加密userID
    cookie := http.Cookie{Name: "userID", Value: userID, Expires: expiration}
    http.SetCookie(c.Writer, &cookie)

    c.JSON(http.StatusOK, gin.H{"msg": "Hello " + userID})
    return
  }else {
    c.JSON(http.StatusBadRequest, gin.H{"msg": "账号或密码错误"})
  }
}

// 检测是否登陆的中间件
func AuthRequired() gin.HandlerFunc {
  return func(c *gin.Context) {
    cookie, _ := c.Request.Cookie("userID")
    if cookie == nil {
      c.JSON(http.StatusUnauthorized, gin.H{"msg": "请先登陆"})
      c.Abort()
    }
    // 实际生产中应校验cookie是否合法
    c.Next()
  }
}

// 查看用户个人信息
func UserInfo(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"msg": "007的个人页面"})
}

// 退出登陆
func Logout(c *gin.Context) {
  // 设置cookie过期
  expiration := time.Now()
  expiration = expiration.AddDate(0, 0, -1)
  cookie := http.Cookie{Name: "userID", Value: "", Expires: expiration}
  http.SetCookie(c.Writer, &cookie)

  c.JSON(http.StatusOK, gin.H{"msg": "退出成功"})
}