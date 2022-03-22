package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisConnection struct {
	Context context.Context
	Client  *redis.Client
}

var RedisConn *redisConnection

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6377",
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
