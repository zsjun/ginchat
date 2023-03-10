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
	// defer db.Close()
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}
	// db.AutoMigrate(&models.Contact{})
	// db.AutoMigrate(&models.GroupBacsic{})

	//user := models.Message{Name: "Alice", PassWord: "123"}
	// result := db.Create(&user)
	// if result.Error != nil {
	// 	t.Fatalf("failed to create user: %v", result.Error)
	// }

}
