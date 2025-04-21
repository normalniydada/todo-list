package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
	"todo-list/config"
)

func ProvideRedisClient(cfg *config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Printf("[ERROR] Failed to connect to redis: %v", err)
		return rdb, err
	}

	log.Println("[INFO] Successfully connected to redis")
	return rdb, nil
}
