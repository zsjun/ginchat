package service

import (
	"fmt"
	"ginchat/common"
	"ginchat/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
func getCurrentUser(c *gin.Context) (userInfo models.UserBasic) {
	session := sessions.Default(c)
	// 类型转换一下
	userInfo = session.Get("currentUser").(models.UserBasic)
	return
}
func setCurrentUser(c *gin.Context, userInfo models.UserBasic) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}
// Login
// @Tags 获取用户列表
// @Param name body  string true "用户名"
// @Param pass_word body string true "密码"
// @Success 200 {json} json{"code","message"}
// @Router /user/Login [post]
func Login(c *gin.Context) {
	session := sessions.Default(c)

	loginUserInfo := models.User{}
	if c.ShouldBindJSON(&loginUserInfo) != nil {
		c.String(http.StatusOK, "参数错误")
		return
	}
	// err := c.ShouldBindJSON(&loginUserInfo)
	// if err != nil {
	// 	fmt.Println(12, err)
	// 	panic(err)
	// }
	fmt.Println(1122, loginUserInfo.Name, loginUserInfo.PassWord)
	
	
	db,err := models.FindUserByName(loginUserInfo.Name)

	if err != nil {
		panic(err)
	}
	// fmt.Println(11, db)
	if loginUserInfo.PassWord == db.PassWord {
		// 邮箱和密码正确则将当前用户信息写入session中
		// setCurrentUser(c, *db)
		c.String(http.StatusOK, "登录成功")
	} else {
		c.String(http.StatusOK, "登录失败")
	}
	// username := c.PostForm("username")
	// password := c.PostForm("password")

	// Validate form input
	// if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
	// 	return
	// }

	// Check for username and password match, usually from a database
	// if username != "hello" || password != "itsme" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	// 	return
	// }

	// Save the username in the session
	// session.Set(common.Userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(common.Userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(common.Userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// me is the handler that will return the user information stored in the
// session.
func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(common.Userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// status is the handler that will tell the user whether it is logged in or not.
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
