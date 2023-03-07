package config

import (
	"ginchat/common"
)

type Mysql struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	PassWord string `json:"password"`
}

func ReadSection(name string) map[string]string {
	return common.VP.GetStringMapString(name)
}

func GetMysqlConfig() (*Mysql, error) {
	mapConfig := ReadSection("mysql")
	mysqlConfig := Mysql{}
	mysqlConfig.Ip = mapConfig["ip"]
	mysqlConfig.Port = mapConfig["port"]
	mysqlConfig.Database = mapConfig["database"]
	mysqlConfig.User = mapConfig["user"]
	mysqlConfig.PassWord = mapConfig["password"]
	return &mysqlConfig, nil
}

type Redis struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

func GetRedisConfig() (*Mysql, error) {
	mapConfig := ReadSection("redis")
	myRedisConfig := Mysql{}
	myRedisConfig.Ip = mapConfig["ip"]
	myRedisConfig.Port = mapConfig["port"]
	return &myRedisConfig, nil
}
