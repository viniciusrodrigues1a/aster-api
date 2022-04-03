package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type redisConnection struct {
	Context context.Context
	Client  *redis.Client
}

var RedisConn *redisConnection

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6375",
		Password: "",
		DB:       0,
	})

	RedisConn = &redisConnection{
		Context: context.Background(),
		Client:  client,
	}
}

func StopRedis() {
	RedisConn.Client.Close()
}
