package handlers

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/services"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// @Summary Создать пользователя
// @Security BearerAuth
// @Description Создает нового пользователя
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(user))
}

// @Summary Получить всех пользователей
// @Security BearerAuth
// @Description Возвращает список всех пользователей
// @Tags Пользователи
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users [get]
func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список пользователей"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(users))
}

// @Summary Получить пользователя
// @Security BearerAuth
// @Description Возвращает пользователя по ID
// @Tags Пользователи
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		if err == repositories.ErrUserNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Пользователь не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении пользователя"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

// @Summary Обновить пользователя
// @Security BearerAuth
// @Description Обновляет данные пользователя
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.User true "Обновленные данные пользователя"
// @Success 200 {object} map[string]interface{} "Пользователь успешно обновлен"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.UserService.UpdateUser(id, &input); err != nil {
		if err == repositories.ErrUserNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Пользователь не найден"))
		} else {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Пользователь успешно обновлен"))
}

// @Summary Удалить пользователя
// @Security BearerAuth
// @Description Удаляет пользователя
// @Tags Пользователи
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]interface{} "Пользователь успешно удален"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	if err := h.UserService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить пользователя"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Пользователь успешно удален"))
}

//// @Summary Аутентификация пользователя
//// @Security BearerAuth
//// @Description Проверяет email или username и пароль для аутентификации
//// @Tags Пользователи
//// @Accept json
//// @Produce json
//// @Param credentials body map[string]string true "Email/Username и пароль"
//// @Success 200 {object} map[string]interface{} "Аутентификация успешна"
//// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
//// @Failure 401 {object} map[string]interface{} "Неверные учетные данные"
//// @Router /users/authenticate [post]
//func (h *UserHandler) AuthenticateUserHandler(c *gin.Context) {
//	var credentials map[string]string
//	if err := c.ShouldBindJSON(&credentials); err != nil {
//		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
//		return
//	}
//
//	identifier := credentials["identifier"]
//	password := credentials["password"]
//
//	if identifier == "" || password == "" {
//		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Обязательные поля: identifier и password"))
//		return
//	}
//
//	user, err := h.UserService.AuthenticateUser(identifier, password)
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
//		return
//	}
//
//	// Здесь можно обновить поле LastLoginAt или сгенерировать JWT токен
//
//	c.JSON(http.StatusOK, utils.SuccessResponse("Аутентификация успешна"))
//}
