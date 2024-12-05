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

type NotificationHandler struct {
	NotificationService services.NotificationService
}

func NewNotificationHandler(notificationService services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		NotificationService: notificationService,
	}
}

// @Summary Создать уведомление
// @Security BearerAuth
// @Description Создает новое уведомление
// @Tags Уведомления
// @Accept json
// @Produce json
// @Param notification body models.Notification true "Данные уведомления"
// @Success 201 {object} models.Notification
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /notifications [post]
func (h *NotificationHandler) CreateNotificationHandler(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.NotificationService.CreateNotification(&notification); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать уведомление"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(notification))
}

// @Summary Получить все уведомления
// @Security BearerAuth
// @Description Возвращает список всех уведомлений
// @Tags Уведомления
// @Produce json
// @Success 200 {array} models.Notification
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /notifications [get]
func (h *NotificationHandler) GetAllNotificationsHandler(c *gin.Context) {
	notifications, err := h.NotificationService.GetAllNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить уведомления"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(notifications))
}

// @Summary Получить уведомление по ID
// @Security BearerAuth
// @Description Возвращает уведомление по его ID
// @Tags Уведомления
// @Produce json
// @Param id path int true "ID уведомления"
// @Success 200 {object} models.Notification
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Уведомление не найдено"
// @Router /notifications/{id} [get]
func (h *NotificationHandler) GetNotificationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID уведомления"))
		return
	}

	notification, err := h.NotificationService.GetNotificationByID(id)
	if err != nil {
		if err == repositories.ErrNotificationNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Уведомление не найдено"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении уведомления"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(notification))
}

// @Summary Обновить уведомление
// @Security BearerAuth
// @Description Обновляет данные уведомления по ID
// @Tags Уведомления
// @Accept json
// @Produce json
// @Param id path int true "ID уведомления"
// @Param notification body models.Notification true "Обновленные данные уведомления"
// @Success 200 {object} models.Notification
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Уведомление не найдено"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /notifications/{id} [put]
func (h *NotificationHandler) UpdateNotificationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID уведомления"))
		return
	}

	var input models.Notification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.NotificationService.UpdateNotification(id, &input); err != nil {
		if err == repositories.ErrNotificationNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Уведомление не найдено"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить уведомление"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Уведомление успешно обновлено"))
}

// @Summary Удалить уведомление
// @Security BearerAuth
// @Description Удаляет уведомление по ID
// @Tags Уведомления
// @Param id path int true "ID уведомления"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /notifications/{id} [delete]
func (h *NotificationHandler) DeleteNotificationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID уведомления"))
		return
	}
	if err := h.NotificationService.DeleteNotification(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить уведомление"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Уведомление успешно удалено"))
}
