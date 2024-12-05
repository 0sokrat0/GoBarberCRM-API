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

type ScheduleHandler struct {
	ScheduleService services.ScheduleService
}

func NewScheduleHandler(scheduleService services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		ScheduleService: scheduleService,
	}
}

// @Summary Создать расписание
// @Security BearerAuth
// @Description Создает новое расписание
// @Tags Расписания
// @Accept json
// @Produce json
// @Param schedule body models.Schedule true "Данные расписания"
// @Success 201 {object} models.Schedule
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules [post]
func (h *ScheduleHandler) CreateScheduleHandler(c *gin.Context) {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if schedule.UserID == 0 || schedule.ScheduleDay == "" || schedule.StartTime == "" || schedule.EndTime == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Все поля обязательны"))
		return
	}

	if err := h.ScheduleService.CreateSchedule(&schedule); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать расписание"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(schedule))
}

// @Summary Получить все расписания
// @Security BearerAuth
// @Description Возвращает список всех расписаний
// @Tags Расписания
// @Produce json
// @Success 200 {array} models.Schedule
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules [get]
func (h *ScheduleHandler) GetAllSchedulesHandler(c *gin.Context) {
	schedules, err := h.ScheduleService.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить расписания"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(schedules))
}

// @Summary Получить расписание
// @Security BearerAuth
// @Description Возвращает данные расписания по ID
// @Tags Расписания
// @Produce json
// @Param id path int true "ID расписания"
// @Success 200 {object} models.Schedule
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Расписание не найдено"
// @Router /schedules/{id} [get]
func (h *ScheduleHandler) GetScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID расписания"))
		return
	}

	schedule, err := h.ScheduleService.GetScheduleByID(id)
	if err != nil {
		if err == repositories.ErrScheduleNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Расписание не найдено"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении расписания"))
		}
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(schedule))
}

// @Summary Обновить расписание
// @Security BearerAuth
// @Description Обновляет данные расписания по ID
// @Tags Расписания
// @Accept json
// @Produce json
// @Param id path int true "ID расписания"
// @Param schedule body models.Schedule true "Обновленные данные расписания"
// @Success 200 {object} models.Schedule
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Расписание не найдено"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules/{id} [put]
func (h *ScheduleHandler) UpdateScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID расписания"))
		return
	}

	var input models.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.ScheduleService.UpdateSchedule(id, &input); err != nil {
		if err == repositories.ErrScheduleNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Расписание не найдено"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить расписание"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Расписание успешно обновлено"))
}

// @Summary Удалить расписание
// @Security BearerAuth
// @Description Удаляет расписание по ID
// @Tags Расписания
// @Param id path int true "ID расписания"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules/{id} [delete]
func (h *ScheduleHandler) DeleteScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID расписания"))
		return
	}

	if err := h.ScheduleService.DeleteSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить расписание"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Расписание успешно удалено"))
}

// @Summary Фильтрация расписаний по пользователю
// @Security BearerAuth
// @Description Возвращает расписания для указанного пользователя
// @Tags Расписания
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Success 200 {array} models.Schedule
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules/filter [get]
func (h *ScheduleHandler) FilterSchedulesByUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	schedules, err := h.ScheduleService.FilterSchedulesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось выполнить фильтрацию расписаний"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(schedules))
}
