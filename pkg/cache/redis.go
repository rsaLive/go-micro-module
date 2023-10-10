package cache

import (
	"github.com/dubbogo/gost/log/logger"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	appConfig "github.com/oa-meeting/pkg/config"
)

var RedisProvider = wire.NewSet(NewRedis)

// TODO 添加连接池

// RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
	//RedisDbName string
)

func NewRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     appConfig.Data.Redis.Addr,
		Password: appConfig.Data.Redis.Password,
		DB:       appConfig.Data.Redis.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		logger.Errorf("connRedis err", err)
		panic(err)
	}
	return RedisClient
}
