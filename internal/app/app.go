package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"todo-list/config"
	"todo-list/internal/api/handlers"
	mw "todo-list/internal/api/middleware"
	"todo-list/internal/api/router"
	"todo-list/internal/domain/service"
	"todo-list/internal/infrastructure/cache/redis"
	"todo-list/internal/infrastructure/database/postgres"
	"todo-list/internal/infrastructure/repository"
)

func Start() {
	// Загрузка конфигурации
	cfg := config.NewConfig()
	log.Println("[INFO] Config loaded successfully", cfg)

	// Инициализации базы данных
	db, err := postgres.ProvideDBClient(&cfg.Database)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize database: %v", err)
	}

	defer func() {
		log.Println("[INFO] Closing database connection...")
		sqlDB, dbErr := db.GetDB().DB()
		if dbErr != nil {
			log.Fatalf("[ERROR] Failed to get underlying *sql.DB for closing: %v", dbErr)
			return
		}
		if closeErr := sqlDB.Close(); closeErr != nil {
			log.Printf("[ERROR] Failed to close database connection: %v", closeErr)
		} else {
			log.Printf("[INFO] Closed database connection successfully")
		}
	}()
	log.Println("[INFO] Database initialized successfully")

	// Инициализация Redis клиента
	redisClient, err := redis.ProvideRedisClient(&cfg.Redis)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize Redis client: %v", err)
	}
	if redisClient == nil {
		defer func() {
			log.Println("[INFO] Closing Redis client connection...")
			if err := redisClient.Close(); err != nil {
				log.Printf("[ERROR] Failed to close Redis client: %v", err)
			} else {
				log.Printf("[INFO] Closed Redis client connection successfully")
			}
		}()
	}

	// Запуск сервера
	e := echo.New()
	e.HideBanner = true

	taskRepo := repository.NewTaskRepository(db.GetDB())
	taskService := service.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	fmt.Println(cfg.RateLimiter)
	if cfg.RateLimiter.Enabled {
		rateLimiterMiddleware := mw.RateLimiterMiddleware(redisClient, &cfg.RateLimiter)
		e.Use(rateLimiterMiddleware)
		log.Println("[INFO] Rate limiter middlware activated")
	} else {
		log.Println("[INFO] Rate limiter disabled in config")
	}

	router.NewRouter(e, taskHandlers)

	serverHTTPAddr := fmt.Sprintf("%s:%d", cfg.Server.HTTP.Host, cfg.Server.HTTP.Port)
	if err = e.Start(serverHTTPAddr); err != nil {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}
