package breaks_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/breaks"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	db.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Миграция для модели Breaks
	err = db.DB.AutoMigrate(&models.Breaks{})
	if err != nil {
		panic("Failed to migrate test database")
	}
}

func teardownTestDB() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func TestCreateBreakHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.POST("/breaks", breaks.CreateBreakHandler)

	// Данные для создания перерыва
	breakInput := models.Breaks{
		UserID:     1,
		BreakStart: time.Now(),
		BreakEnd:   time.Now().Add(30 * time.Minute),
	}

	body, _ := json.Marshal(breakInput)

	req := httptest.NewRequest(http.MethodPost, "/breaks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверка стандартизированного ответа
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(breakInput.UserID), data["user_id"])
}

func TestGetAllBreaksHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	breaksData := []models.Breaks{
		{UserID: 1, BreakStart: time.Now(), BreakEnd: time.Now().Add(15 * time.Minute)},
		{UserID: 2, BreakStart: time.Now().Add(1 * time.Hour), BreakEnd: time.Now().Add(1*time.Hour + 30*time.Minute)},
	}
	db.DB.Create(&breaksData)

	router := gin.Default()
	router.GET("/breaks", breaks.GetAllBreaksHandler)

	req := httptest.NewRequest(http.MethodGet, "/breaks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверка данных
	assert.Equal(t, true, response["success"])
	data := response["data"].([]interface{})
	assert.Equal(t, 2, len(data))
}

func TestGetBreakHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	breakData := models.Breaks{
		UserID:     1,
		BreakStart: time.Now(),
		BreakEnd:   time.Now().Add(15 * time.Minute),
	}
	db.DB.Create(&breakData)

	router := gin.Default()
	router.GET("/breaks/:id", breaks.GetBreakHandler)

	req := httptest.NewRequest(http.MethodGet, "/breaks/"+strconv.Itoa(int(breakData.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверка данных
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(breakData.UserID), data["user_id"])
}

func TestUpdateBreakHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	breakData := models.Breaks{
		UserID:     1,
		BreakStart: time.Now(),
		BreakEnd:   time.Now(),
	}
	db.DB.Create(&breakData)

	router := gin.Default()
	router.PUT("/breaks/:id", breaks.UpdateBreakHandler)

	// Новые данные для обновления
	updatedBreak := models.Breaks{
		UserID:     2,
		BreakStart: time.Now(),
		BreakEnd:   time.Now().Add(1 * time.Hour),
	}
	body, _ := json.Marshal(updatedBreak)

	req := httptest.NewRequest(http.MethodPut, "/breaks/"+strconv.Itoa(int(breakData.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверка данных
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(updatedBreak.UserID), data["user_id"])
}

func TestDeleteBreakHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	breakData := models.Breaks{
		UserID:     1,
		BreakStart: time.Now(),
		BreakEnd:   time.Now().Add(15 * time.Minute),
	}
	db.DB.Create(&breakData)

	router := gin.Default()
	router.DELETE("/breaks/:id", breaks.DeleteBreakHandler)

	req := httptest.NewRequest(http.MethodDelete, "/breaks/"+strconv.Itoa(int(breakData.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверка стандартизированного ответа
	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Перерыв успешно удалён", response["data"])
}
