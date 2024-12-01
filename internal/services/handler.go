package services

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Создать услугу
// @Description Создает новую услугу
// @Tags Услуги
// @Accept json
// @Produce json
// @Param service body models.Services true "Данные услуги"
// @Success 201 {object} models.Services
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services [post]
func CreateServiceHandler(c *gin.Context) {
	var service models.Services
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if service.Price <= 0 || service.Duration <= 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Цена и продолжительность должны быть больше нуля"))
		return
	}

	if err := db.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать услугу"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(service))
}

// @Summary Получить все услуги
// @Description Возвращает список всех услуг
// @Tags Услуги
// @Produce json
// @Success 200 {array} models.Services
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services [get]
func GetAllServicesHandler(c *gin.Context) {
	var services []models.Services
	if err := db.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список услуг"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(services))
}

// @Summary Получить услугу
// @Description Возвращает данные услуги по ID
// @Tags Услуги
// @Produce json
// @Param id path int true "ID услуги"
// @Success 200 {object} models.Services
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Услуга не найдена"
// @Router /services/{id} [get]
func GetServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	var service models.Services
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Услуга не найдена"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(service))
}

// @Summary Обновить услугу
// @Description Обновляет данные услуги по ID
// @Tags Услуги
// @Accept json
// @Produce json
// @Param id path int true "ID услуги"
// @Param service body models.Services true "Обновленные данные услуги"
// @Success 200 {object} models.Services
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Услуга не найдена"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services/{id} [put]
func UpdateServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	var service models.Services
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Услуга не найдена"))
		return
	}

	var input models.Services
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if input.Price <= 0 || input.Duration <= 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Цена и продолжительность должны быть больше нуля"))
		return
	}

	if err := db.DB.Model(&service).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить услугу"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(service))
}

// @Summary Удалить услугу
// @Description Удаляет услугу по ID
// @Tags Услуги
// @Param id path int true "ID услуги"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services/{id} [delete]
func DeleteServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	if err := db.DB.Delete(&models.Services{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить услугу"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Услуга успешно удалена"))
}

// @Summary Деактивировать услугу
// @Description Помечает услугу как неактивную
// @Tags Услуги
// @Param id path int true "ID услуги"
// @Success 200 {object} map[string]interface{} "Сообщение об успешной деактивации"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /services/{id}/deactivate [put]
func DeactivateServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	if err := db.DB.Model(&models.Services{}).Where("id = ?", id).Update("is_active", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось деактивировать услугу"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Услуга успешно деактивирована"))
}
