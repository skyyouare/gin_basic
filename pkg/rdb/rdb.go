package rdb

import (
	"context"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/setting"

	"github.com/redis/go-redis/v9"
)

var Conn *redis.Client

func Setup() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.IPHost,
		Password: setting.RedisSetting.PassWord,
		DB:       setting.RedisSetting.Db,
		Protocol: setting.RedisSetting.Protocol,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(err)
	}
	Conn = rdb
}

func Close() {
	defer Conn.Close()
}
