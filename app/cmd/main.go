package main

import (
	"fmt"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/app"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db"
	"log"

	"github.com/0sokrat0/GoGRAFFApi.git/app/configs"
	_ "github.com/0sokrat0/GoGRAFFApi.git/app/docs"
)

// @title GoGRAFF API
// @version 1.0
// @description API для управления GoGRAFF
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	cfg, err := configs.LoadConfig("/root/app/configs")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SslMode,
	)

	err = db.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	router := app.SetupRouter(db.DB)

	err = router.Run(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
