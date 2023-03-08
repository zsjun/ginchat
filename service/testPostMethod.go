package service

import (
	"fmt"
	"ginchat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestPostMethod(c *gin.Context) {
	var user models.User
	// send api field should match json
	// for example
	// use
	// type User struct {
	// 	Name     string `gorm:"unique" json:"name"`
	// 	PassWord string `json:"pass_word"`
	// }
	// send api field shoule is pass_word
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Do something with the user data, e.g. save it to a database
	fmt.Println(22, user.PassWord, user)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %s : %s", user.Name, user.PassWord),
	})

}
