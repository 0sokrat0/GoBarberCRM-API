package main

import (
	"fmt"
	"log"

	"github.com/0sokrat0/GoGRAFFApi.git/configs"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/routes"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Формируем строку подключения к базе данных
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SslMode,
	)

	log.Printf("Connecting to database: %s", dsn)

	// Инициализируем базу данных
	err = db.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Настраиваем маршрутизатор
	router := routes.SetupRouter()

	// Запускаем сервер
	log.Printf("Starting %s API in %s mode on port %d", cfg.App.Name, cfg.App.Environment, cfg.App.Port)
	log.Printf("Server is running at http://localhost:%d", cfg.App.Port)
	err = router.Run(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
