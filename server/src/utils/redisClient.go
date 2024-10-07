package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type contextKey string

const RedisClientKey contextKey = "redisClient"

func InitRedisClient(ctx context.Context) (context.Context, error) {
	url := os.Getenv("redis_url")
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(opts)

	ctx = context.WithValue(ctx, RedisClientKey, redisClient)
	return ctx, nil
}

func GetRedisClientFromContext(ctx context.Context) (*redis.Client, error) {
	redisClient, ok := ctx.Value(RedisClientKey).(*redis.Client)
	if !ok || redisClient == nil {
		return nil, fmt.Errorf("redis client not found in context")
	}
	return redisClient, nil
}
