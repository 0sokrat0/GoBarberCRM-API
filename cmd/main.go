package main

import (
	"fmt"
	"log"

	"github.com/0sokrat0/GoGRAFFApi.git/configs"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
)

func main() {

	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SslMode)

	log.Printf("Connecting to database: %s", dsn)

	log.Printf("Starting %s in %s mode on port %d", cfg.App.Name, cfg.App.Environment, cfg.App.Port)

	err = db.InitDB(dsn)
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %v", err)
	}
}
