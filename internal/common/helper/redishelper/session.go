package redishelper

import (
	"context"
	"example/internal/common/helper/redisclienthelper"
	"fmt"
	"time"
)

type (
	RedisSessionHelper interface {
		Exists(ctx context.Context, key string) error
		Get(ctx context.Context, key string, value interface{}) error
		GetInterface(ctx context.Context, key string, value interface{}) (interface{}, error)
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Del(ctx context.Context, key string) error
		Expire(ctx context.Context, key string, expiration time.Duration) error
		SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	}
)

func GenerateRedisSessionKey(resource string, userId string) string {
	return fmt.Sprintf("session:%s#%s", resource, userId)
}

func NewRedisSessionHelper(redisClientHelper *redisclienthelper.RedisClientHelper) RedisSessionHelper {
	return &redisHelper{
		client: redisClientHelper.Client,
	}
}
