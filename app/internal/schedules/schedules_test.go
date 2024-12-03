package schedules_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/schedules"
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

	err = db.DB.AutoMigrate(&models.Schedules{})
	if err != nil {
		panic("Failed to migrate test database")
	}
}

func teardownTestDB() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func TestCreateScheduleHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.POST("/schedules", schedules.CreateScheduleHandler)

	schedule := models.Schedules{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "17:00",
	}
	body, _ := json.Marshal(schedule)

	req := httptest.NewRequest(http.MethodPost, "/schedules", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(schedule.UserID), data["user_id"])
	assert.Equal(t, schedule.ScheduleDay, data["schedule_day"])
}

func TestGetAllSchedulesHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	schedulesData := []models.Schedules{
		{UserID: 1, ScheduleDay: "Monday", StartTime: "09:00", EndTime: "17:00"},
		{UserID: 2, ScheduleDay: "Tuesday", StartTime: "10:00", EndTime: "18:00"},
	}
	db.DB.Create(&schedulesData)

	router := gin.Default()
	router.GET("/schedules", schedules.GetAllSchedulesHandler)

	req := httptest.NewRequest(http.MethodGet, "/schedules", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].([]interface{})
	assert.Equal(t, 2, len(data))
}

func TestGetScheduleHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	schedule := models.Schedules{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "17:00",
	}
	db.DB.Create(&schedule)

	router := gin.Default()
	router.GET("/schedules/:id", schedules.GetScheduleHandler)

	req := httptest.NewRequest(http.MethodGet, "/schedules/"+strconv.Itoa(int(schedule.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(schedule.UserID), data["user_id"])
}

func TestUpdateScheduleHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	schedule := models.Schedules{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "17:00",
	}
	db.DB.Create(&schedule)

	router := gin.Default()
	router.PUT("/schedules/:id", schedules.UpdateScheduleHandler)

	updatedSchedule := models.Schedules{
		UserID:      2,
		ScheduleDay: "Tuesday",
		StartTime:   "10:00",
		EndTime:     "18:00",
	}
	body, _ := json.Marshal(updatedSchedule)

	req := httptest.NewRequest(http.MethodPut, "/schedules/"+strconv.Itoa(int(schedule.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(updatedSchedule.UserID), data["user_id"])
}

func TestDeleteScheduleHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	schedule := models.Schedules{
		UserID:      1,
		ScheduleDay: "Monday",
		StartTime:   "09:00",
		EndTime:     "17:00",
	}
	db.DB.Create(&schedule)

	router := gin.Default()
	router.DELETE("/schedules/:id", schedules.DeleteScheduleHandler)

	req := httptest.NewRequest(http.MethodDelete, "/schedules/"+strconv.Itoa(int(schedule.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Расписание успешно удалено", response["data"])
}

func TestFilterSchedulesByUserHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создание тестовых данных
	schedulesData := []models.Schedules{
		{UserID: 1, ScheduleDay: "Monday", StartTime: "09:00", EndTime: "17:00"},
		{UserID: 1, ScheduleDay: "Tuesday", StartTime: "10:00", EndTime: "18:00"},
		{UserID: 2, ScheduleDay: "Wednesday", StartTime: "11:00", EndTime: "19:00"},
	}
	db.DB.Create(&schedulesData)

	router := gin.Default()
	router.GET("/schedules/filter", schedules.FilterSchedulesByUserHandler)

	req := httptest.NewRequest(http.MethodGet, "/schedules/filter?user_id=1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].([]interface{})
	assert.Equal(t, 2, len(data))
}
