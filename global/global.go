package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// Upper case var can export
	VP        *viper.Viper
	DB        *gorm.DB
	Secret    []byte
	Red       *redis.Client
	WsClients map[*websocket.Conn]bool

	// clients stores all connected WebSocket clients
)

const (
	Userkey string = "user"
)

var JwtKey = []byte("my_secret_key")

// const (
// 	userkey string = "user"
// )
