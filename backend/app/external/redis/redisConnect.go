package redis

import (
	"fmt"
	"github.com/Bottlist/bottlist/pkg/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisConnector() *redis.Client {
	conf := config.LoadConfig()

	conn := redisConnInfo(*conf.RedisInfo)

	return conn
}

func redisConnInfo(redisInfo config.RedisInfo) *redis.Client {

	addr := fmt.Sprintf("%s:%s",
		redisInfo.RedisHost,
		redisInfo.RedisPort,
	)
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return conn
}
