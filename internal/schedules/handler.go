package schedules

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Создать расписание
// @Description Создает новое расписание
// @Tags Расписания
// @Accept json
// @Produce json
// @Param schedule body models.Schedules true "Данные расписания"
// @Success 201 {object} models.Schedules
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules [post]
func CreateScheduleHandler(c *gin.Context) {
	var schedule models.Schedules
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if schedule.UserID == 0 || schedule.ScheduleDay == "" || schedule.StartTime == "" || schedule.EndTime == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Все поля обязательны"))
		return
	}

	if err := db.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать расписание"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(schedule))
}

// @Summary Получить все расписания
// @Description Возвращает список всех расписаний
// @Tags Расписания
// @Produce json
// @Success 200 {array} models.Schedules
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules [get]
func GetAllSchedulesHandler(c *gin.Context) {
	var schedules []models.Schedules
	if err := db.DB.Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить расписания"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(schedules))
}

// @Summary Получить расписание
// @Description Возвращает данные расписания по ID
// @Tags Расписания
// @Produce json
// @Param id path int true "ID расписания"
// @Success 200 {object} models.Schedules
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Расписание не найдено"
// @Router /schedules/{id} [get]
func GetScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID расписания"))
		return
	}

	var schedule models.Schedules
	if err := db.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Расписание не найдено"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(schedule))
}

// @Summary Обновить расписание
// @Description Обновляет данные расписания по ID
// @Tags Расписания
// @Accept json
// @Produce json
// @Param id path int true "ID расписания"
// @Param schedule body models.Schedules true "Обновленные данные расписания"
// @Success 200 {object} models.Schedules
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Расписание не найдено"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules/{id} [put]
func UpdateScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID расписания"))
		return
	}

	var schedule models.Schedules
	if err := db.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Расписание не найдено"))
		return
	}

	var input models.Schedules
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := db.DB.Model(&schedule).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить расписание"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(schedule))
}

// @Summary Удалить расписание
// @Description Удаляет расписание по ID
// @Tags Расписания
// @Param id path int true "ID расписания"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules/{id} [delete]
func DeleteScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID расписания"))
		return
	}

	if err := db.DB.Delete(&models.Schedules{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить расписание"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Расписание успешно удалено"))
}

// @Summary Фильтрация расписаний по пользователю
// @Description Возвращает расписания для указанного пользователя
// @Tags Расписания
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Success 200 {array} models.Schedules
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /schedules/filter [get]
func FilterSchedulesByUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	var schedules []models.Schedules
	if err := db.DB.Where("user_id = ?", userID).Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось выполнить фильтрацию расписаний"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(schedules))
}
