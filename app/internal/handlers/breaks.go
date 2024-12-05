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

type BreakHandler struct {
	BreakService services.BreakService
}

func NewBreakHandler(breakService services.BreakService) *BreakHandler {
	return &BreakHandler{
		BreakService: breakService,
	}
}

// @Summary Создать перерыв
// @Description Создает новый перерыв
// @Security BearerAuth
// @Tags Перерывы
// @Accept json
// @Produce json
// @Param break body models.Break true "Данные перерыва"
// @Success 201 {object} models.Break
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks [post]
func (h *BreakHandler) CreateBreakHandler(c *gin.Context) {
	var breakModel models.Break
	if err := c.ShouldBindJSON(&breakModel); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.BreakService.CreateBreak(&breakModel); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать перерыв"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(breakModel))
}

// @Summary Получить все перерывы
// @Security BearerAuth
// @Description Возвращает список всех перерывов
// @Tags Перерывы
// @Produce json
// @Success 200 {array} models.Break
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks [get]
func (h *BreakHandler) GetAllBreaksHandler(c *gin.Context) {
	breaks, err := h.BreakService.GetAllBreaks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список перерывов"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(breaks))
}

// @Summary Получить перерыв
// @Security BearerAuth
// @Description Возвращает данные перерыва по ID
// @Tags Перерывы
// @Produce json
// @Param id path int true "ID перерыва"
// @Success 200 {object} models.Break
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Перерыв не найден"
// @Router /breaks/{id} [get]
func (h *BreakHandler) GetBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID перерыва"))
		return
	}

	breakModel, err := h.BreakService.GetBreakByID(id)
	if err != nil {
		if err == repositories.ErrBreakNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Перерыв не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении перерыва"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(breakModel))
}

// @Summary Обновить перерыв
// @Security BearerAuth
// @Description Обновляет данные перерыва по ID
// @Tags Перерывы
// @Accept json
// @Produce json
// @Param id path int true "ID перерыва"
// @Param break body models.Break true "Обновленные данные перерыва"
// @Success 200 {object} models.Break
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Перерыв не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks/{id} [put]
func (h *BreakHandler) UpdateBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID перерыва"))
		return
	}

	var input models.Break
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.BreakService.UpdateBreak(id, &input); err != nil {
		if err == repositories.ErrBreakNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Перерыв не найден"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить перерыв"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Перерыв успешно обновлён"))
}

// @Summary Удалить перерыв
// @Security BearerAuth
// @Description Удаляет перерыв по ID
// @Tags Перерывы
// @Param id path int true "ID перерыва"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /breaks/{id} [delete]
func (h *BreakHandler) DeleteBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID перерыва"))
		return
	}

	if err := h.BreakService.DeleteBreak(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить перерыв"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Перерыв успешно удалён"))
}
