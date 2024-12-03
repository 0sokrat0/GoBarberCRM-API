package main

import (
	"fmt"
	"log"

	"github.com/0sokrat0/GoGRAFFApi.git/app/configs"
	_ "github.com/0sokrat0/GoGRAFFApi.git/app/docs"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/routes"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db"
)

// @title GoGRAFF API
// @version 1.0
// @description API для управления GoGRAFF
// @host localhost:8080
// @BasePath /
func main() {

	cfg, err := configs.LoadConfig("./app/configs")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SslMode,
	)
	log.Printf("Connecting to database: %s", dsn)

	err = db.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	router := routes.SetupRouter()

	// Запускаем сервер
	log.Printf("Starting %s API in %s mode on port %d", cfg.App.Name, cfg.App.Environment, cfg.App.Port)
	log.Printf("Server is running at http://localhost:%d", cfg.App.Port)
	err = router.Run(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
