package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T, models ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		t.Fatalf("Не удалось выполнить миграцию базы данных: %v", err)
	}

	return db
}
