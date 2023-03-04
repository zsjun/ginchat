package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/asaskevich/govalidator"
)

// GetUserList
// @Tags 获取用户列表
// @Success 200 {json} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	// userList := make([]*models.UserBasic, 10)
	userList, err := models.UserBasic{}.List()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"message": userList,
	})
}


// CreateUser
// @Tags 创建用户
// @Success 200 {json} json{"code","message"}
// @Param name body  string true "用户名"
// @Param pass_word body string true "密码"
// @Router /user/create [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "参数错误",
		})
		return
	}
	_, err = govalidator.ValidateStruct(user)
	name := c.Query("name")
	data := models.FindUserByName(name)
	if data != nil {
		c.JSON(200, gin.H{
			"message": "用户名已经注册",
		})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{
			"message": "用户名和密码不符合标准",
		})
		return
	}

	err = models.UserBasic{}.Create(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "创建成功",
	})
	return
}

// UpdateUser
// @Tags 更新用户
// @Success 200 {json} json{"code","message"}
// @Param id body int true "id"
// @Param name body string true "用户名"
// @Param pass_word body string true "密码"
// @Router /user/Update [put]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "参数错误",
		})
		return
	}
	err = models.UserBasic{}.Update(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "创建成功",
	})
	return
}

// DeleteUser
// @Tags 删除用户
// @Param id path  int true "用户信息"
// @Success 200 {json} json{"code","message"}
// @Router /user/delete/{id} [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	err := c.ShouldBind(&user)
	t, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(t)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "参数错误",
		})
		return
	}
	err = user.Delete()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
	return 
}