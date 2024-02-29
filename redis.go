package main

import (
	"context"
	"time"
)

func get(ctx context.Context, key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

func set(ctx context.Context, key, value string, expiration time.Duration) error {
	err := redisClient.Set(ctx, key, value, expiration).Err()
	return err
}

func incr(ctx context.Context, key string, expiration time.Duration) (val int64, err error) {
	val, err = redisClient.Incr(ctx, key).Result()
	if err != nil {
		return
	}

	_, err = redisClient.Expire(ctx, key, expiration).Result()
	if err != nil {
		return
	}

	return
}
