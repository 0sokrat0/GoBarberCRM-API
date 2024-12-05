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

type ServiceHandler struct {
	ServiceService services.ServiceService
}

func NewServiceHandler(serviceService services.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		ServiceService: serviceService,
	}
}

// @Summary Создать услугу
// @Security BearerAuth
// @Description Создает новую услугу
// @Tags Услуги
// @Accept json
// @Produce json
// @Param service body models.Service true "Данные услуги"
// @Success 201 {object} models.Service
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services [post]
func (h *ServiceHandler) CreateServiceHandler(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if service.Price <= 0 || service.Duration <= 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Цена и продолжительность должны быть больше нуля"))
		return
	}

	if err := h.ServiceService.CreateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать услугу"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(service))
}

// @Summary Получить все услуги
// @Security BearerAuth
// @Description Возвращает список всех услуг
// @Tags Услуги
// @Produce json
// @Success 200 {array} models.Service
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services [get]
func (h *ServiceHandler) GetAllServicesHandler(c *gin.Context) {
	services, err := h.ServiceService.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список услуг"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(services))
}

// @Summary Получить услугу
// @Security BearerAuth
// @Description Возвращает данные услуги по ID
// @Tags Услуги
// @Produce json
// @Param id path int true "ID услуги"
// @Success 200 {object} models.Service
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Услуга не найдена"
// @Router /services/{id} [get]
func (h *ServiceHandler) GetServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	service, err := h.ServiceService.GetServiceByID(id)
	if err != nil {
		if err == repositories.ErrServiceNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Услуга не найдена"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении услуги"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(service))
}

// @Summary Обновить услугу
// @Security BearerAuth
// @Description Обновляет данные услуги по ID
// @Tags Услуги
// @Accept json
// @Produce json
// @Param id path int true "ID услуги"
// @Param service body models.Service true "Обновленные данные услуги"
// @Success 200 {object} models.Service
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Услуга не найдена"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services/{id} [put]
func (h *ServiceHandler) UpdateServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	var input models.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if input.Price <= 0 || input.Duration <= 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Цена и продолжительность должны быть больше нуля"))
		return
	}

	if err := h.ServiceService.UpdateService(id, &input); err != nil {
		if err == repositories.ErrServiceNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Услуга не найдена"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить услугу"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Услуга успешно обновлена"))
}

// @Summary Удалить услугу
// @Security BearerAuth
// @Description Удаляет услугу по ID
// @Tags Услуги
// @Param id path int true "ID услуги"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services/{id} [delete]
func (h *ServiceHandler) DeleteServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	if err := h.ServiceService.DeleteService(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить услугу"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Услуга успешно удалена"))
}

// @Summary Деактивировать услугу
// @Security BearerAuth
// @Description Помечает услугу как неактивную
// @Tags Услуги
// @Param id path int true "ID услуги"
// @Success 200 {object} map[string]interface{} "Сообщение об успешной деактивации"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services/{id}/deactivate [put]
func (h *ServiceHandler) DeactivateServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	if err := h.ServiceService.DeactivateService(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось деактивировать услугу"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Услуга успешно деактивирована"))
}
