package db

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error

	// Открываем соединение с базой данных
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Выполняем миграции
	err = DB.AutoMigrate(
		&models.Bookings{},
		&models.User{},
		&models.Client{},
		&models.Schedule{},
		&models.Service{},
		&models.HistoryLogs{},
		&models.Notification{},
		&models.Break{},
		&models.AuthUser{},
	)
	if err != nil {
		return err
	}

	log.Println("Database connection established and migrations applied successfully.")
	return nil
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
