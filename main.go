package main

import (
	"ginchat/router"
	"ginchat/utils"
)

// 不操作数据库，把所有用户信息写死在代码里
// var db = &User{Id: 10001, Email: "abc@gmail.cn", Username: "Alice", Password: "123456"}

func main() {
	// utils.Init()
	// 注册User结构体
	// gob.Register(User{})
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := router.Router()

	// This handler will match /user/john but will not match /user/ or /user

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
