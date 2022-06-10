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
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	RedisConn = &redisConnection{
		Context: ctx,
		Client:  client,
	}
}

func StopRedis() {
	RedisConn.Client.Close()
}
