package main

import (
	"encoding/gob"
	"fmt"
	"ginchat/router"
	"ginchat/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func getCurrentUser(c *gin.Context) (userInfo User) {
	session := sessions.Default(c)
  // 类型转换一下
	userInfo = session.Get("currentUser").(User) 
	return
}

func setCurrentUser(c *gin.Context, userInfo User) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
  // 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save() 
}

func setupRouter(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		loginUserInfo := User{}
		if c.ShouldBindJSON(&loginUserInfo) != nil {
			c.String(http.StatusOK, "参数错误")
			return
		}
		if loginUserInfo.Email == db.Email && loginUserInfo.Password == db.Password {
      // 邮箱和密码正确则将当前用户信息写入session中
			setCurrentUser(c, *db)
			c.String(http.StatusOK, "登录成功")
		} else {
			c.String(http.StatusOK, "登录失败")
		}
	})

	r.GET("/sayHello", func(c *gin.Context) {
		userInfo := getCurrentUser(c)
		c.String(http.StatusOK, "Hello "+userInfo.Username)
	})
}
// 不操作数据库，把所有用户信息写死在代码里
var db = &User{Id: 10001, Email: "abc@gmail.cn", Username: "Alice", Password: "123456"} 



func main() {
  // 注册User结构体
  gob.Register(User{}) 
  utils.InitConfig()
  utils.InitMysql()
  r := router.Router()
  fmt.Println("hello, world1")
  // 设置生成sessionId的密钥
  store := cookie.NewStore([]byte("secret")) 
  // mysession是返回給前端的sessionId名
  r.Use(sessions.Sessions("mysession", store))

  setupRouter(r)

  

  r.GET("/hello", func(c *gin.Context) {
    session := sessions.Default(c)

    if session.Get("hello") != "world" {
      session.Set("hello", "world")
      session.Save()
    }

    c.JSON(200, gin.H{"hello": session.Get("hello")})
  })


  // 需要登陆保护的
  // auth := r.Group("/api")
	// auth.Use(AuthRequired())
	// {
	// 	auth.GET("/me", UserInfo)
	// 	auth.GET("/logout", Logout)
	// }
  r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 登陆
// func Login(c *gin.Context) {
//   // 为了演示方便，我直接通过url明文传递账号密码，实际生产中应该用HTTP POST在body中传递
//   userID := c.Query("user_id")
//   password := c.Query("password")
//   fmt.Println(userID,password)
//   // 用户身份校验（查询数据库）
//   if userID == "007" && password == "007" {
//     // 生成cookie
//     expiration := time.Now()
//     expiration = expiration.AddDate(0, 0, 1)
//     // 实际生产中我们可以加密userID
//     cookie := http.Cookie{Name: "userID", Value: userID, Expires: expiration}
//     http.SetCookie(c.Writer, &cookie)

//     c.JSON(http.StatusOK, gin.H{"msg": "Hello " + userID})
//     return
//   }else {
//     c.JSON(http.StatusBadRequest, gin.H{"msg": "账号或密码错误"})
//   }
// }

// // 检测是否登陆的中间件
// func AuthRequired() gin.HandlerFunc {
//   return func(c *gin.Context) {
//     cookie, _ := c.Request.Cookie("userID")
//     if cookie == nil {
//       c.JSON(http.StatusUnauthorized, gin.H{"msg": "请先登陆"})
//       c.Abort()
//     }
//     // 实际生产中应校验cookie是否合法
//     c.Next()
//   }
// }

// // 查看用户个人信息
// func UserInfo(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{"msg": "007的个人页面"})
// }

// // 退出登陆
// func Logout(c *gin.Context) {
//   // 设置cookie过期
//   expiration := time.Now()
//   expiration = expiration.AddDate(0, 0, -1)
//   cookie := http.Cookie{Name: "userID", Value: "", Expires: expiration}
//   http.SetCookie(c.Writer, &cookie)

//   c.JSON(http.StatusOK, gin.H{"msg": "退出成功"})
// }