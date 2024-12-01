package clients

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Создать клиента
// @Description Создает нового клиента
// @Tags Клиенты
// @Accept json
// @Produce json
// @Param client body models.Clients true "Данные клиента"
// @Success 201 {object} models.Clients
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients [post]
func CreateClientHandler(c *gin.Context) {
	var client models.Clients

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := db.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать клиента"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(client))
}

// @Summary Получить всех клиентов
// @Description Возвращает список всех клиентов
// @Tags Клиенты
// @Produce json
// @Success 200 {array} models.Clients
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients [get]
func GetAllClientsHandler(c *gin.Context) {
	var clients []models.Clients

	if err := db.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список клиентов"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(clients))
}

// @Summary Получить клиента
// @Description Возвращает данные клиента по ID
// @Tags Клиенты
// @Produce json
// @Param id path int true "ID клиента"
// @Success 200 {object} models.Clients
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Router /clients/{id} [get]
func GetClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}
	var client models.Clients
	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент не найден"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(client))
}

// @Summary Обновить клиента
// @Description Обновляет данные клиента по ID
// @Tags Клиенты
// @Accept json
// @Produce json
// @Param id path int true "ID клиента"
// @Param client body models.Clients true "Обновленные данные клиента"
// @Success 200 {object} models.Clients
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/{id} [put]
func UpdateClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}

	var client models.Clients
	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент не найден"))
		return
	}

	var input models.Clients
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := db.DB.Model(&client).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить клиента"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(client))
}

// @Summary Удалить клиента
// @Description Удаляет клиента по ID
// @Tags Клиенты
// @Param id path int true "ID клиента"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/{id} [delete]
func DeleteClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}
	if err := db.DB.Delete(&models.Clients{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить клиента"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Клиент успешно удалён"))
}

// @Summary Получить клиента по Telegram ID
// @Description Возвращает данные клиента по Telegram ID
// @Tags Клиенты
// @Produce json
// @Param tg_id path int true "Telegram ID клиента"
// @Success 200 {object} models.Clients
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Router /clients/telegram/{tg_id} [get]
func GetClientByTelegramIDHandler(c *gin.Context) {
	tgID := c.Param("tg_id")

	var client models.Clients
	if err := db.DB.Where("tg_id = ?", tgID).First(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент с таким Telegram ID не найден"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(client))
}

// @Summary Фильтрация клиентов по имени
// @Description Возвращает список клиентов по имени
// @Tags Клиенты
// @Produce json
// @Param name query string true "Имя клиента для фильтрации"
// @Success 200 {array} models.Clients
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/filter [get]
func FilterClientsByNameHandler(c *gin.Context) {
	name := c.Query("name")

	var clients []models.Clients
	if err := db.DB.Where("LOWER(first_name) LIKE LOWER(?) OR LOWER(last_name) LIKE LOWER(?)", "%"+name+"%", "%"+name+"%").Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось выполнить фильтрацию клиентов"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(clients))
}

// @Summary Быстрое добавление клиента
// @Description Создает клиента с минимальными данными (например, через Telegram и/или номер телефона)
// @Tags Клиенты
// @Accept json
// @Produce json
// @Param client body models.Clients true "Минимальные данные клиента"
// @Success 201 {object} models.Clients
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/quick_add [post]
func QuickAddClientHandler(c *gin.Context) {
	var client models.Clients
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	// Проверяем, что хотя бы одно из обязательных полей указано
	if client.PhoneNumber == "" && client.TgID == 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Номер телефона или Telegram ID обязательны"))
		return
	}

	// Проверяем на существование клиента по номеру телефона или Telegram ID
	var existingClient models.Clients
	if client.PhoneNumber != "" {
		if err := db.DB.Where("phone_number = ?", client.PhoneNumber).First(&existingClient).Error; err == nil {
			c.JSON(http.StatusConflict, utils.ErrorResponse("Клиент с таким номером телефона уже существует"))
			return
		}
	}
	if client.TgID != 0 {
		if err := db.DB.Where("tg_id = ?", client.TgID).First(&existingClient).Error; err == nil {
			c.JSON(http.StatusConflict, utils.ErrorResponse("Клиент с таким Telegram ID уже существует"))
			return
		}
	}

	// Создаём нового клиента
	if err := db.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось добавить клиента"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(client))
}

// @Summary Найти клиента по контактным данным
// @Description Ищет клиента по Email или номеру телефона
// @Tags Клиенты
// @Produce json
// @Param email query string false "Email клиента"
// @Param phone query string false "Номер телефона клиента"
// @Success 200 {object} models.Clients
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Router /clients/search [get]
func SearchClientHandler(c *gin.Context) {
	email := c.Query("email")
	phone := c.Query("phone")

	var client models.Clients
	if email != "" {
		if err := db.DB.Where("email = ?", email).First(&client).Error; err == nil {
			c.JSON(http.StatusOK, utils.SuccessResponse(client))
			return
		}
	}

	if phone != "" {
		if err := db.DB.Where("phone_number = ?", phone).First(&client).Error; err == nil {
			c.JSON(http.StatusOK, utils.SuccessResponse(client))
			return
		}
	}

	c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент с такими контактными данными не найден"))
}

// @Summary Проверить существование клиента
// @Description Проверяет, существует ли клиент с заданными контактными данными
// @Tags Клиенты
// @Produce json
// @Param phone_number query string false "Номер телефона клиента"
// @Param tg_id query int false "Telegram ID клиента"
// @Success 200 {object} map[string]bool "Результат проверки"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Router /clients/check [get]
func CheckClientExistenceHandler(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	tgID := c.Query("tg_id")

	var count int64
	if phoneNumber != "" {
		db.DB.Model(&models.Clients{}).Where("phone_number = ?", phoneNumber).Count(&count)
	} else if tgID != "" {
		db.DB.Model(&models.Clients{}).Where("tg_id = ?", tgID).Count(&count)
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(map[string]bool{"exists": count > 0}))
}
