package main

import (
	"ginchat/router"
	"ginchat/utils"
	"log"
	"os/exec"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	exec.Command("dir").Run()
	err := utils.InitConfig()
	if err != nil {
		log.Fatalf("Init fail err: %s", err)
		return
	}
	utils.InitMysql()
}
func main() {
	r := router.Router()
	r.Run()
}
