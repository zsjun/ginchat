package test

import (
	"ginchat/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateTableName(t *testing.T) {
	dsn := "root:guilai123@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			t.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {	
			t.Fatalf("failed to migrate database: %v", err)
	}

	user := models.UserBasic{Name: "Alice", PassWord: "123"}
	result := db.Create(&user)
	if result.Error != nil {
			t.Fatalf("failed to create user: %v", result.Error)
	}

}