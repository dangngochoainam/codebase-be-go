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

type Resource string

const (
	ACCESS_TOKEN  Resource = "access_token"
	REFRESH_TOKEN Resource = "refresh_token"
)

func GenerateRedisSessionKey(resource Resource, Id string) string {
	return fmt.Sprintf("session:%s#%s", resource, Id)
}

func NewRedisSessionHelper(redisClientHelper *redisclienthelper.RedisClientHelper) RedisSessionHelper {
	return &redisHelper{
		client: redisClientHelper.Client,
	}
}
