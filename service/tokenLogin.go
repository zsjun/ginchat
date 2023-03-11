package service

import (
	"fmt"
	"ginchat/global"
	"ginchat/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// func getCurrentUser(c *gin.Context) (userInfo models.UserBasic) {
// 	session := sessions.Default(c)
// 	// 类型转换一下
// 	userInfo = session.Get("currentUser").(models.UserBasic)
// 	return
// }
// func setCurrentUser(c *gin.Context, userInfo models.UserBasic) {
// 	session := sessions.Default(c)
// 	session.Set("currentUser", userInfo)
// 	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
// 	session.Save()
// }

// Login
// @Tags 获取用户列表
// @Param name body  string true "用户名"
// @Param pass_word body string true "密码"
// @Success 200 {json} json{"code","message"}
// @Router /user/Login [post]
func TokenLogin(c *gin.Context) {
	loginUserInfo := models.UserBasic{}
	if c.ShouldBindJSON(&loginUserInfo) != nil {
		c.String(http.StatusOK, "参数错误")
		return
	}
	db, err := models.FindUserByName(loginUserInfo.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	if loginUserInfo.PassWord == db.PassWord {
		// Create JWT token
		expirationTime := time.Now().Add(24 * time.Hour)
		loginUserInfo = models.UserBasic{
			UserName: loginUserInfo.UserName,
			UserId:   loginUserInfo.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, loginUserInfo)
		tokenString, err := token.SignedString(global.JwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// c.JSON(http.StatusOK, gin.H{"token": tokenString})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "login success", "token": tokenString})
	} else {
		c.String(http.StatusOK, "登录失败")
		return
	}

}

func TokenLogOut(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Mark the token as expired
	claims := &models.UserBasic{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return global.JwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims.ExpiresAt = time.Now().Unix() - 1 // Mark token as expired
	token.Claims = claims
	_, err = token.SignedString(global.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
