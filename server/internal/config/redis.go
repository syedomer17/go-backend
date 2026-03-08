package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func init() {
	connectRedis()
}

func connectRedis() {
	cfg, err := Load()

	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.UPSTASH_REDIS_REST_URL,
		Password: cfg.UPSTASH_REDIS_REST_TOKEN,
	})

	_, err = RedisClient.Ping(ctx).Result()

	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}
