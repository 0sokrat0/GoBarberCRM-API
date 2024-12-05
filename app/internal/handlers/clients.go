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

type ClientHandler struct {
	ClientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) *ClientHandler {
	return &ClientHandler{
		ClientService: clientService,
	}
}

// @Summary Создать клиента
// @Security BearerAuth
// @Description Создает нового клиента
// @Tags Клиенты
// @Accept json
// @Produce json
// @Param client body models.Client true "Данные клиента"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients [post]
func (h *ClientHandler) CreateClientHandler(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.ClientService.CreateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать клиента"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(client))
}

// @Summary Получить всех клиентов
// @Security BearerAuth
// @Description Возвращает список всех клиентов
// @Tags Клиенты
// @Produce json
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients [get]
func (h *ClientHandler) GetAllClientsHandler(c *gin.Context) {
	clients, err := h.ClientService.GetAllClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список клиентов"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(clients))
}

// @Summary Получить клиента
// @Security BearerAuth
// @Description Возвращает данные клиента по ID
// @Tags Клиенты
// @Produce json
// @Param id path int true "ID клиента"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Router /clients/{id} [get]
func (h *ClientHandler) GetClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}

	client, err := h.ClientService.GetClientByID(id)
	if err != nil {
		if err == repositories.ErrClientNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении клиента"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(client))
}

// @Summary Обновить клиента
// @Security BearerAuth
// @Description Обновляет данные клиента по ID
// @Tags Клиенты
// @Accept json
// @Produce json
// @Param id path int true "ID клиента"
// @Param client body models.Client true "Обновленные данные клиента"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/{id} [put]
func (h *ClientHandler) UpdateClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}

	var input models.Client
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.ClientService.UpdateClient(id, &input); err != nil {
		if err == repositories.ErrClientNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить клиента"))
		}
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Клиент успешно обновлён"))
}

// @Summary Удалить клиента
// @Security BearerAuth
// @Description Удаляет клиента по ID
// @Tags Клиенты
// @Param id path int true "ID клиента"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/{id} [delete]
func (h *ClientHandler) DeleteClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}
	if err := h.ClientService.DeleteClient(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить клиента"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Клиент успешно удалён"))
}

// @Summary Получить клиента по Telegram ID
// @Security BearerAuth
// @Description Возвращает данные клиента по Telegram ID
// @Tags Клиенты
// @Produce json
// @Param tg_id path int true "Telegram ID клиента"
// @Success 200 {object} models.Client
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Router /clients/telegram/{tg_id} [get]
func (h *ClientHandler) GetClientByTelegramIDHandler(c *gin.Context) {
	tgIDStr := c.Param("tg_id")
	tgID, err := strconv.ParseInt(tgIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный Telegram ID"))
		return
	}

	client, err := h.ClientService.GetClientByTelegramID(tgID)
	if err != nil {
		if err == repositories.ErrClientNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент с таким Telegram ID не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении клиента"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(client))
}

// @Summary Фильтрация клиентов по имени
// @Security BearerAuth
// @Description Возвращает список клиентов по имени
// @Tags Клиенты
// @Produce json
// @Param name query string true "Имя клиента для фильтрации"
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/filter [get]
func (h *ClientHandler) FilterClientsByNameHandler(c *gin.Context) {
	name := c.Query("name")

	clients, err := h.ClientService.FilterClientsByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось выполнить фильтрацию клиентов"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(clients))
}

// @Summary Быстрое добавление клиента
// @Security BearerAuth
// @Description Создает клиента с минимальными данными (например, через Telegram и/или номер телефона)
// @Tags Клиенты
// @Accept json
// @Produce json
// @Param client body models.Client true "Минимальные данные клиента"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 409 {object} map[string]interface{} "Клиент уже существует"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /clients/quick_add [post]
func (h *ClientHandler) QuickAddClientHandler(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.ClientService.QuickAddClient(&client); err != nil {
		if err.Error() == "номер телефона или Telegram ID обязательны" {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		} else if err == repositories.ErrClientAlreadyExists {
			c.JSON(http.StatusConflict, utils.ErrorResponse("Клиент уже существует"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось добавить клиента"))
		}
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(client))
}

// @Summary Найти клиента по контактным данным
// @Security BearerAuth
// @Description Ищет клиента по Email или номеру телефона
// @Tags Клиенты
// @Produce json
// @Param email query string false "Email клиента"
// @Param phone query string false "Номер телефона клиента"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Клиент не найден"
// @Router /clients/search [get]
func (h *ClientHandler) SearchClientHandler(c *gin.Context) {
	email := c.Query("email")
	phone := c.Query("phone")

	if email == "" && phone == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Необходимо указать email или номер телефона"))
		return
	}

	client, err := h.ClientService.SearchClientByEmailOrPhone(email, phone)
	if err != nil {
		if err == repositories.ErrClientNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Клиент с такими контактными данными не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при поиске клиента"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(client))
}

// @Summary Проверить существование клиента
// @Security BearerAuth
// @Description Проверяет, существует ли клиент с заданными контактными данными
// @Tags Клиенты
// @Produce json
// @Param phone_number query string false "Номер телефона клиента"
// @Param tg_id query int false "Telegram ID клиента"
// @Success 200 {object} map[string]bool "Результат проверки"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Router /clients/check [get]
func (h *ClientHandler) CheckClientExistenceHandler(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	tgIDStr := c.Query("tg_id")

	var tgID int64
	var err error
	if tgIDStr != "" {
		tgID, err = strconv.ParseInt(tgIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный tg_id"))
			return
		}
	}

	exists, err := h.ClientService.CheckClientExistence(phoneNumber, tgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(map[string]bool{"exists": exists}))
}
