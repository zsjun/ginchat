package main

import (
	"fmt"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
  utils.Init()
  utils.InitConfig()
  utils.InitMysql()
  r := router.Router()
  fmt.Println("hello, world1")
  r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}