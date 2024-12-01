package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Создать пользователя
// @Description Создает нового пользователя
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param user body models.Users true "Данные пользователя"
// @Success 201 {object} models.Users
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}
	log.Println("Received user data:", user)

	if user.Username == "" || user.PasswordHash == "" || user.Role == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Обязательные поля: Username, Password и Role"))
		return
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось хэшировать пароль"))
		return
	}
	user.PasswordHash = string(hashedPassword)

	// Сохранение в базу данных
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать пользователя"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(user))
}

// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags Пользователи
// @Produce json
// @Success 200 {array} models.Users
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users [get]
func GetAllUsersHandler(c *gin.Context) {
	var users []models.Users
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список пользователей"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(users))
}

// @Summary Получить пользователя
// @Description Возвращает пользователя по ID
// @Tags Пользователи
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.Users
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	var user models.Users
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Пользователь не найден"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

// @Summary Обновить пользователя
// @Description Обновляет данные пользователя
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.Users true "Обновленные данные пользователя"
// @Success 200 {object} models.Users
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	var user models.Users
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Пользователь не найден"))
		return
	}

	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if input.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось хешировать пароль"))
			return
		}
		input.PasswordHash = string(hashedPassword)
	}

	if err := db.DB.Model(&user).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить пользователя"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

// @Summary Удалить пользователя
// @Description Удаляет пользователя
// @Tags Пользователи
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	if err := db.DB.Delete(&models.Users{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить пользователя"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Пользователь успешно удалён"))
}

// @Summary Аутентификация пользователя
// @Description Проверяет email и пароль для аутентификации
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Email и пароль"
// @Success 200 {object} map[string]interface{} "Сообщение об успешной аутентификации"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 401 {object} map[string]interface{} "Неверные учетные данные"
// @Router /users/authenticate [post]
func AuthenticateUserHandler(c *gin.Context) {
	var credentials map[string]string
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	email := credentials["email"]
	password := credentials["password"]

	var user models.Users
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Неверные учетные данные"))
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Неверные учетные данные"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Аутентификация успешна"))
}
