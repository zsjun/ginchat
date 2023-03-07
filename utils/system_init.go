package utils

import (
	"context"
	"fmt"
	"ginchat/common"
	"ginchat/config"
	"ginchat/models"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() error {
	common.VP = viper.GetViper()
	common.VP.SetConfigName("app")
	common.VP.AddConfigPath("config")
	return common.VP.ReadInConfig()
}

func InitMysql() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	},
	)
	mysqlConfig, err := config.GetMysqlConfig()
	if err != nil {
		panic("failed to read mysqlConfig")
	}

	common.DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		mysqlConfig.User, mysqlConfig.PassWord,
		mysqlConfig.Ip, mysqlConfig.Port, mysqlConfig.Database)),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	common.DB.AutoMigrate(&models.UserBasic{})
}
func InitRedis() {
	myRedisConfig, err := config.GetRedisConfig()
	if err != nil {
		panic("failed to read mysqlConfig")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     myRedisConfig.Ip + ":" + myRedisConfig.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(1123)
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(1112233)
		panic(err)
	}
	fmt.Println("222", pong)
}

// func Init() {
// 	// common.userkey = "user";
// 	fmt.Println(common)
// }
