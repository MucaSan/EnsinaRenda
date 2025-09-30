package redis

import (
	"context"
	context_config "ensina-renda/config/context"

	"github.com/redis/go-redis/v9"
)

const RedisContextKey context_config.ContextKey = "redis"

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func GetRedis(ctx context.Context) *redis.Client {
	return ctx.Value(RedisContextKey).(*redis.Client)
}
