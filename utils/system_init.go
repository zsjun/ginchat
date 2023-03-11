package utils

import (
	"context"
	"fmt"
	"ginchat/config"
	"ginchat/global"
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
	global.VP = viper.GetViper()
	global.VP.SetConfigName("app")
	global.VP.AddConfigPath("config")
	return global.VP.ReadInConfig()
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

	global.DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		mysqlConfig.User, mysqlConfig.PassWord,
		mysqlConfig.Ip, mysqlConfig.Port, mysqlConfig.Database)),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	global.DB.AutoMigrate(&models.UserBasic{})
}
func InitRedis() {
	myRedisConfig, err := config.GetRedisConfig()
	if err != nil {
		panic("failed to read mysqlConfig")
	}
	global.Red = redis.NewClient(&redis.Options{
		Addr:     myRedisConfig.Ip + ":" + myRedisConfig.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	_, err = global.Red.Ping(ctx).Result()
	if err != nil {

		panic(err)
	}
}

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, msg string) error {
	err := global.Red.Publish(ctx, channel, msg)
	if err != nil {
		fmt.Println("Redis Publish fail")
		panic(err)
	}
	return nil
}

func Subscribel(ctx context.Context, channel string) (string, error) {
	sub := global.Red.PSubscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println("Redis Subscribel fail")
		panic(err)
	}
	return msg.Payload, err
}
