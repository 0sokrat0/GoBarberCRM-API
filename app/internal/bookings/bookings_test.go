package bookings_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/bookings"
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

	err = db.DB.AutoMigrate(&models.Bookings{})
	if err != nil {
		panic("Failed to migrate test database")
	}
}

func teardownTestDB() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func TestCreateBookingHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.POST("/bookings", bookings.CreateBookingHandler)

	// Данные для создания бронирования
	booking := models.Bookings{
		UserID:      1,
		ClientID:    1,
		ServiceID:   1,
		BookingTime: time.Now(),
	}

	body, _ := json.Marshal(booking)

	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверяем стандартизированный формат ответа
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(1), data["user_id"])
	assert.Equal(t, float64(1), data["client_id"])
	assert.Equal(t, float64(1), data["service_id"])
}

func TestCreateBookingHandler_Duplicate(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.POST("/bookings", bookings.CreateBookingHandler)

	// Создаём дубликат бронирования
	booking := models.Bookings{
		UserID:      1,
		ClientID:    1,
		ServiceID:   1,
		BookingTime: time.Now(),
	}
	db.DB.Create(&booking)

	body, _ := json.Marshal(booking)

	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, false, response["success"])
	assert.Contains(t, response["error"], "Временной слот уже занят")
}

func TestGetBookingHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.GET("/bookings/:id", bookings.GetBookingHandler)

	// Создаём бронирование
	booking := models.Bookings{
		UserID:      1,
		ClientID:    1,
		ServiceID:   1,
		BookingTime: time.Now(),
	}
	db.DB.Create(&booking)

	req := httptest.NewRequest(http.MethodGet, "/bookings/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверяем стандартизированный формат ответа
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(1), data["user_id"])
}

func TestGetBookingHandler_NotFound(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.GET("/bookings/:id", bookings.GetBookingHandler)

	req := httptest.NewRequest(http.MethodGet, "/bookings/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, false, response["success"])
	assert.Contains(t, response["error"], "Бронирование не найдено")
}

func TestUpdateBookingHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.PUT("/bookings/:id", bookings.UpdateBookingHandler)

	// Создаём бронирование
	booking := models.Bookings{
		UserID:      1,
		ClientID:    1,
		ServiceID:   1,
		BookingTime: time.Now(),
	}
	db.DB.Create(&booking)

	// Обновляем бронирование
	updatedBooking := models.Bookings{
		ServiceID: 2,
	}
	body, _ := json.Marshal(updatedBooking)

	req := httptest.NewRequest(http.MethodPut, "/bookings/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверяем стандартизированный формат ответа
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(2), data["service_id"])
}

func TestDeleteBookingHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.DELETE("/bookings/:id", bookings.DeleteBookingHandler)

	// Создаём бронирование
	booking := models.Bookings{
		UserID:      1,
		ClientID:    1,
		ServiceID:   1,
		BookingTime: time.Now(),
	}
	db.DB.Create(&booking)

	req := httptest.NewRequest(http.MethodDelete, "/bookings/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверяем стандартизированный формат ответа
	assert.Equal(t, true, response["success"])
	assert.Contains(t, response["data"], "Бронирование успешно удалено")
}
