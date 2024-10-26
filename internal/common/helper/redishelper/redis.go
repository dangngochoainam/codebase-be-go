package redishelper

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type (
	redisHelper struct {
		client *redis.Client
	}
)

func (h *redisHelper) Exists(ctx context.Context, key string) (err error) {
	indicator, err := h.client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if indicator == 0 {
		return redis.Nil
	}
	return nil
}

func (h *redisHelper) Get(ctx context.Context, key string, value interface{}) (err error) {
	data, err := h.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), &value)
	if err != nil {
		return err
	}
	return nil
}

func (h *redisHelper) GetInterface(ctx context.Context, key string, value interface{}) (interface{}, error) {
	var err error
	data, err := h.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	typeValue := reflect.TypeOf(value)
	kind := typeValue.Kind()

	var outData interface{}
	switch kind {
	case reflect.Ptr, reflect.Struct, reflect.Slice:
		outData = reflect.New(typeValue).Interface()
	default:
		outData = reflect.Zero(typeValue).Interface()
	}
	err = json.Unmarshal([]byte(data), &outData)
	if err != nil {
		return nil, err
	}

	switch kind {
	case reflect.Ptr, reflect.Struct, reflect.Slice:
		outDataValue := reflect.ValueOf(outData)

		if reflect.Indirect(reflect.ValueOf(outDataValue)).IsZero() {
			return nil, errors.New("Get redis nill result")
		}
		if outDataValue.IsZero() {
			return outDataValue.Interface(), nil
		}
		return outDataValue.Elem().Interface(), nil
	}
	var outValue interface{} = outData
	if reflect.TypeOf(outData).ConvertibleTo(typeValue) {
		outValueConverted := reflect.ValueOf(outData).Convert(typeValue)
		outValue = outValueConverted.Interface()
	}
	return outValue, nil
}

func (h *redisHelper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = h.client.Set(ctx, key, string(data), expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (h *redisHelper) Del(ctx context.Context, key string) (err error) {
	_, err = h.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (h *redisHelper) Expire(ctx context.Context, key string, expiration time.Duration) (err error) {
	_, err = h.client.Expire(ctx, key, expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (h *redisHelper) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (isSucces bool, err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	isSucces, err = h.client.SetNX(ctx, key, string(data), expiration).Result()
	if err != nil {
		return false, err
	}
	return isSucces, nil
}
