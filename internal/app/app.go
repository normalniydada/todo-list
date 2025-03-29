package app

import (
	"log"
	"todo-list/config"
	"todo-list/internal/infrastructure/database/postgres"
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

	//// Запуск сервера
	//router := gin.Default()
}
