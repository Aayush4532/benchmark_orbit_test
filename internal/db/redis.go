package db

import (
	"Orbit/configs"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func SetUpRedis() {
	cfg := configs.LoadConfig().Redis

	var options *redis.Options

	if strings.Contains(cfg.Url, "://") {

		parsed, err := redis.ParseURL(cfg.Url)
		if err != nil {
			log.Fatalf("invalid REDIS_URL: %v", err)
		}

		options = parsed

	} else {

		options = &redis.Options{
			Addr: cfg.Url,
			DB:   0,
		}
	}

	options.PoolSize = 64
	options.MinIdleConns = 8

	// Timeouts
	options.PoolTimeout = 15 * time.Second
	options.ReadTimeout = 15 * time.Second
	options.WriteTimeout = 15 * time.Second

	redisClient = redis.NewClient(options)
}

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		SetUpRedis()
	})
	return redisClient
}