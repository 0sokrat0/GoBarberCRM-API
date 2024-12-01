package services_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/0sokrat0/GoGRAFFApi.git/internal/services"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
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
	err = db.DB.AutoMigrate(&models.Services{})
	if err != nil {
		panic("Failed to migrate test database")
	}
}

func teardownTestDB() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func TestCreateServiceHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.POST("/services", services.CreateServiceHandler)

	service := models.Services{
		Name:        "Test Service",
		Description: "Test Description",
		Price:       100.0,
		Duration:    60,
	}

	body, _ := json.Marshal(service)
	req := httptest.NewRequest(http.MethodPost, "/services", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, service.Name, data["name"])
	assert.Equal(t, service.Description, data["description"])
	assert.Equal(t, service.Price, data["price"].(float64))
	assert.Equal(t, service.Duration, int(data["duration"].(float64)))
}

func TestGetAllServicesHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	db.DB.Create(&models.Services{
		Name:        "Service 1",
		Description: "Description 1",
		Price:       50.0,
		Duration:    30,
	})
	db.DB.Create(&models.Services{
		Name:        "Service 2",
		Description: "Description 2",
		Price:       150.0,
		Duration:    90,
	})

	router := gin.Default()
	router.GET("/services", services.GetAllServicesHandler)

	req := httptest.NewRequest(http.MethodGet, "/services", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].([]interface{})
	assert.Len(t, data, 2)
}

func TestGetServiceHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	service := models.Services{
		Name:        "Service 1",
		Description: "Description 1",
		Price:       50.0,
		Duration:    30,
	}
	db.DB.Create(&service)

	router := gin.Default()
	router.GET("/services/:id", services.GetServiceHandler)

	req := httptest.NewRequest(http.MethodGet, "/services/"+strconv.Itoa(int(service.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, service.Name, data["name"])
	assert.Equal(t, service.Description, data["description"])
}

func TestUpdateServiceHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	service := models.Services{
		Name:        "Service 1",
		Description: "Description 1",
		Price:       50.0,
		Duration:    30,
	}
	db.DB.Create(&service)

	router := gin.Default()
	router.PUT("/services/:id", services.UpdateServiceHandler)

	updatedService := models.Services{
		Name:        "Updated Service",
		Description: "Updated Description",
		Price:       100.0,
		Duration:    60,
	}

	body, _ := json.Marshal(updatedService)
	req := httptest.NewRequest(http.MethodPut, "/services/"+strconv.Itoa(int(service.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, updatedService.Name, data["name"])
	assert.Equal(t, updatedService.Description, data["description"])
}

func TestDeleteServiceHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	service := models.Services{
		Name:        "Service 1",
		Description: "Description 1",
		Price:       50.0,
		Duration:    30,
	}
	db.DB.Create(&service)

	router := gin.Default()
	router.DELETE("/services/:id", services.DeleteServiceHandler)

	req := httptest.NewRequest(http.MethodDelete, "/services/"+strconv.Itoa(int(service.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Услуга успешно удалена", response["data"])
}
