package users_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/users"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	db.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	err = db.DB.AutoMigrate(&models.Users{})
	if err != nil {
		panic("Failed to migrate test database")
	}
}

func teardownTestDB() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func TestCreateUserHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := gin.Default()
	router.POST("/users", users.CreateUserHandler)

	user := models.Users{
		Username:     "testuser",
		PasswordHash: "mypassword", // Пароль будет хэшироваться
		Role:         "admin",
		Email:        "test@example.com",
		PhoneNumber:  "+1234567890",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Ошибка: ожидаемый статус 201, получен %d. Ответ сервера: %s", w.Code, w.Body.String())
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Ошибка парсинга ответа")
	assert.Equal(t, true, response["success"], "Ответ должен быть успешным")

	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok, "Поле 'data' отсутствует или имеет неверный тип")
	assert.Equal(t, user.Username, data["username"], "Некорректное имя пользователя")

	storedUser := models.Users{}
	err = db.DB.Where("username = ?", user.Username).First(&storedUser).Error
	assert.NoError(t, err, "Ошибка получения пользователя из базы")

	// Проверка хэша пароля
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(user.PasswordHash))
	assert.NoError(t, err, "Пароль не совпадает с хэшированным значением")
}

func TestGetAllUsersHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создаем тестовые данные
	usersData := []models.Users{
		{Username: "user1", PasswordHash: "hash1", Role: "user", Email: "user1@example.com", PhoneNumber: "+1234567891"},
		{Username: "user2", PasswordHash: "hash2", Role: "admin", Email: "user2@example.com", PhoneNumber: "+1234567892"},
	}
	db.DB.Create(&usersData)

	router := gin.Default()
	router.GET("/users", users.GetAllUsersHandler)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
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

func TestGetUserHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	user := models.Users{
		Username:     "testuser",
		PasswordHash: "hash",
		Role:         "user",
		Email:        "test@example.com",
		PhoneNumber:  "+1234567890",
	}
	db.DB.Create(&user)

	router := gin.Default()
	router.GET("/users/:id", users.GetUserHandler)

	req := httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, user.Username, data["username"])
	assert.Equal(t, user.Role, data["role"])
	assert.Equal(t, user.Email, data["email"])
}

func TestUpdateUserHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	user := models.Users{
		Username:     "olduser",
		PasswordHash: "oldhash",
		Role:         "user",
		Email:        "old@example.com",
		PhoneNumber:  "+1234567890",
	}
	db.DB.Create(&user)

	router := gin.Default()
	router.PUT("/users/:id", users.UpdateUserHandler)

	updatedUser := models.Users{
		Username: "newuser",
		Role:     "admin",
		Email:    "new@example.com",
	}
	body, _ := json.Marshal(updatedUser)

	req := httptest.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, updatedUser.Username, data["username"])
	assert.Equal(t, updatedUser.Role, data["role"])
	assert.Equal(t, updatedUser.Email, data["email"])
}

func TestDeleteUserHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	user := models.Users{
		Username:     "testuser",
		PasswordHash: "hash",
		Role:         "user",
		Email:        "test@example.com",
		PhoneNumber:  "+1234567890",
	}
	db.DB.Create(&user)

	router := gin.Default()
	router.DELETE("/users/:id", users.DeleteUserHandler)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Пользователь успешно удалён", response["data"])
}

func TestPasswordHashing(t *testing.T) {
	// Исходный пароль
	plainPassword := "mypassword"

	// Хэшируем пароль
	hash, err := utils.HashPassword(plainPassword)
	assert.NoError(t, err, "Ошибка хэширования пароля")
	assert.NotEmpty(t, hash, "Хэш пароля не должен быть пустым")

	// Проверяем, что хэш не совпадает с исходным паролем
	assert.NotEqual(t, plainPassword, hash, "Хэш пароля не должен совпадать с оригиналом")

	// Проверяем соответствие пароля хэшу
	match := utils.CheckPasswordHash(plainPassword, hash)
	assert.True(t, match, "Пароль не соответствует хэшу")

	// Проверяем несоответствие неверного пароля
	wrongPassword := "wrongpassword"
	match = utils.CheckPasswordHash(wrongPassword, hash)
	assert.False(t, match, "Неверный пароль не должен соответствовать хэшу")
}

func TestAuthenticateUserHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Создаем пользователя
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.Users{
		Username:     "testuser",
		PasswordHash: string(hashedPassword),
		Email:        "test@example.com",
	}
	db.DB.Create(&user)

	router := gin.Default()
	router.POST("/users/authenticate", users.AuthenticateUserHandler)

	// Тестируем успешную аутентификацию
	credentials := map[string]string{"email": "test@example.com", "password": "password123"}
	body, _ := json.Marshal(credentials)

	req := httptest.NewRequest(http.MethodPost, "/users/authenticate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Ожидался статус 200")
}
