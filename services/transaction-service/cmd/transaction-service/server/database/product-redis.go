package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ProductRedisConn *redisConnection

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})

	ProductRedisConn = &redisConnection{
		Context: context.Background(),
		Client:  client,
	}
}

func StopProductRedis() {
	ProductRedisConn.Client.Close()
}
