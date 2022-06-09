package statestorelib

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type RedisStateStoreRepository struct {
	Context context.Context
	Client  *redis.Client
}

func New(context context.Context, client *redis.Client) *RedisStateStoreRepository {
	return &RedisStateStoreRepository{
		Context: context,
		Client:  client,
	}
}

type StateStoreWriter interface {
	StoreState(id string, state interface{}) error
}

func (r *RedisStateStoreRepository) StoreState(id string, state interface{}) error {
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return err
	}

	redisErr := r.Client.Set(r.Context, id, stateJSON, 0).Err()
	if redisErr != nil {
		return redisErr
	}

	return nil
}

type StateStoreReader interface {
	ReadState(id string) (string, error)
}

func (r *RedisStateStoreRepository) ReadState(id string) (string, error) {
	val, err := r.Client.Get(r.Context, id).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
