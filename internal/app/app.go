package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"todo-list/config"
	"todo-list/internal/api/handlers"
	"todo-list/internal/api/router"
	"todo-list/internal/domain/service"
	"todo-list/internal/infrastructure/database/postgres"
	"todo-list/internal/infrastructure/repository"
)

func Start() {
	// Загрузка конфигурации
	cfg := config.NewConfig()
	log.Println("[INFO] Config loaded successfully", cfg)

	// Инициализации базы данных
	db, err := postgres.ProvideDB(&cfg.Database)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize database: %v", err)
	}
	defer db.Close()

	log.Println("[INFO] Database initialized successfully")

	// Запуск сервера
	e := echo.New()

	taskRepo := repository.NewTaskRepository(db.GetDB())
	taskService := service.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))

	router.NewRouter(e, taskHandlers)

	serverHTTPAddr := fmt.Sprintf("%s:%d", cfg.Server.HTTP.Host, cfg.Server.HTTP.Port)
	if err = e.Start(serverHTTPAddr); err != nil {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}
