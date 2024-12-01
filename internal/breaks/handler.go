package breaks

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Создать перерыв
// @Description Создает новый перерыв
// @Tags Перерывы
// @Accept json
// @Produce json
// @Param break body models.Breaks true "Данные перерыва"
// @Success 201 {object} models.Breaks
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks [post]
func CreateBreakHandler(c *gin.Context) {
	var breaks models.Breaks
	if err := c.ShouldBindJSON(&breaks); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := db.DB.Create(&breaks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать перерыв"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(breaks))
}

// @Summary Получить все перерывы
// @Description Возвращает список всех перерывов
// @Tags Перерывы
// @Produce json
// @Success 200 {array} models.Breaks
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks [get]
func GetAllBreaksHandler(c *gin.Context) {
	var breaks []models.Breaks
	if err := db.DB.Find(&breaks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список перерывов"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(breaks))
}

// @Summary Получить перерыв
// @Description Возвращает данные перерыва по ID
// @Tags Перерывы
// @Produce json
// @Param id path int true "ID перерыва"
// @Success 200 {object} models.Breaks
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Перерыв не найден"
// @Router /breaks/{id} [get]
func GetBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID перерыва"))
		return
	}

	var breaks models.Breaks
	if err := db.DB.First(&breaks, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Перерыв не найден"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(breaks))
}

// @Summary Обновить перерыв
// @Description Обновляет данные перерыва по ID
// @Tags Перерывы
// @Accept json
// @Produce json
// @Param id path int true "ID перерыва"
// @Param break body models.Breaks true "Обновленные данные перерыва"
// @Success 200 {object} models.Breaks
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Перерыв не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks/{id} [put]
func UpdateBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID перерыва"))
		return
	}

	var breaks models.Breaks
	if err := db.DB.First(&breaks, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Перерыв не найден"))
		return
	}

	var input models.Breaks
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := db.DB.Model(&breaks).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить перерыв"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(breaks))
}

// @Summary Удалить перерыв
// @Description Удаляет перерыв по ID
// @Tags Перерывы
// @Param id path int true "ID перерыва"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks/{id} [delete]
func DeleteBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID перерыва"))
		return
	}

	if err := db.DB.Delete(&models.Breaks{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить перерыв"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Перерыв успешно удалён"))
}
