package main

func RedisCfg() {

	if Init_Redis {

		FolderCheck("dao/redis", "dao/redis", "[REDIS] ")
		WriteContentToConfigYaml(Redis_Init_Content, "dao/redis/redis.go", "[REDIS] ")
		WriteContentToConfigYaml(Redis_Config_Yaml, "config.yaml", "[REDIS] ")
	}
}

var (
	Redis_Config_Yaml = `redis: 
  host: "127.0.0.1"
  password: ""
  port: 6379
  db: 0
  pool_size: 100`
	Redis_Init_Content = `package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: fmt.Sprintf("%s", viper.GetString("redis.password")),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("close Mysql DB failed, err: %v\n ", zap.Error(err))
	}
}`
)
