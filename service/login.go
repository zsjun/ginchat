package service

import (
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
	loginUserInfo := models.User{}
	if c.ShouldBindJSON(&loginUserInfo) != nil {
		c.String(http.StatusOK, "参数错误")
		return
	}
	db, err := models.FindUserByName(loginUserInfo.Name)
	if err != nil {
		panic(err)
	}
	if loginUserInfo.PassWord == db.PassWord {
		// 邮箱和密码正确则将当前用户信息写入session中
		// setCurrentUser(c, *db)
		// c.String(http.StatusOK, "登录成功")
		session := sessions.Default(c)
		session.Set(common.Userkey, db.ID)
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "login success"})
	} else {
		c.String(http.StatusOK, "登录失败")
		return
	}

}

func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(common.Userkey)
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(common.Userkey)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// me is the handler that will return the user information stored in the
// session.
func IsLogin(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(common.Userkey)
	if userID == nil {
		// If not authenticated, redirect to login
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userID, "status": "You are logged in"})
}
