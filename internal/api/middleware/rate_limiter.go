package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
	"todo-list/config"
)

func RateLimiterMiddleware(redisClient *redis.Client, cfg *config.RateLimiterConfig) echo.MiddlewareFunc {
	if !cfg.Enabled {
		log.Println("[INFO] Rate limiter disabled in config")
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return next(c)
			}
		}
	}

	log.Printf("[INFO] Rate limiter enabled: limit=%d requests / %v", cfg.Limit, cfg.Window)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			if ip == "" {
				log.Println("[WARN] Rate limiter: Could not get client IP address.")
				return next(c)
			}

			redisKey := fmt.Sprintf("rate_limit_%s", ip)

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			var count int64
			pipe := redisClient.TxPipeline()
			incrCmd := pipe.Incr(ctx, redisKey)
			pipe.Expire(ctx, redisKey, cfg.Window)

			_, err := pipe.Exec(ctx)

			if err != nil {
				log.Printf("[ERROR] Rate limiter: Redis error for key %s: %v", redisKey, err)
				return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Could not process request rate limit"})
			}

			count, err = incrCmd.Result()
			if err != nil {
				log.Printf("[ERROR] Rate limiter: Could not get INCR result for key %s: %v", redisKey, err)
				return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Could not process request rate limit count"})
			}

			if count > int64(cfg.Limit) {
				log.Printf("[INFO] Rate limit exceeded for IP: %s (Count: %d, Limit: %d)", ip, count, cfg.Limit)

				c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", cfg.WindowSec))
				c.Response().Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", cfg.Limit))
				c.Response().Header().Set("X-RateLimit-Remaining", "0")
				resetTime := time.Now().Add(cfg.Window).Unix()
				c.Response().Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", resetTime))

				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": cfg.ErrorMessage})
			}

			remaining := int64(cfg.Limit) - count
			if remaining < 0 {
				remaining = 0
			}
			c.Response().Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", cfg.Limit))
			c.Response().Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
			resetTime := time.Now().Add(cfg.Window).Unix()
			c.Response().Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", resetTime))

			return next(c)
		}
	}

}
