package redisclienthelper

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	RedisClientHelper struct {
		Client *redis.Client
	}
	RedisConfigOptions struct {
		Address  string
		Password string
		DB       int
	}
)

func NewClientRedisHelper(opts *RedisConfigOptions) *RedisClientHelper {
	client := redis.NewClient(&redis.Options{
		Addr:     opts.Address,
		Password: opts.Password,
		DB:       opts.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.S().Panic("Failed to init redis", zap.Error(err))
	}
	return &RedisClientHelper{
		Client: client,
	}

}
