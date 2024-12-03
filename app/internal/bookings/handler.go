package bookings

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Создать бронирование
// @Description Создает новое бронирование, если слот времени не занят
// @Tags Бронирования
// @Accept json
// @Produce json
// @Param booking body models.Bookings true "Данные бронирования"
// @Success 201 {object} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 409 {object} map[string]interface{} "Слот времени уже занят"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings [post]
func CreateBookingHandler(c *gin.Context) {
	var booking models.Bookings

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	var count int64
	db.DB.Model(&models.Bookings{}).
		Where("user_id = ? AND booking_time = ?", booking.UserID, booking.BookingTime).
		Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, utils.ErrorResponse("Временной слот уже занят"))
		return
	}

	if err := db.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать бронирование"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(booking))
}

// @Summary Получить бронирование
// @Description Получает бронирование по ID
// @Tags Бронирования
// @Produce json
// @Param id path int true "ID бронирования"
// @Success 200 {object} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID"
// @Failure 404 {object} map[string]interface{} "Бронирование не найдено"
// @Router /bookings/{id} [get]
func GetBookingHandler(c *gin.Context) {
	var booking models.Bookings
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}

	if err := db.DB.Preload("Client").Preload("Service").Preload("User").First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Бронирование не найдено"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(booking))
}

// @Summary Получить все бронирования
// @Description Получает список всех бронирований
// @Tags Бронирования
// @Produce json
// @Success 200 {array} models.Bookings
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings [get]
func GetAllBookingsHandler(c *gin.Context) {
	var bookings []models.Bookings
	if err := db.DB.Preload("Client").Preload("Service").Preload("User").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список бронирований"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(bookings))
}

// @Summary Обновить бронирование
// @Description Обновляет данные бронирования по ID
// @Tags Бронирования
// @Accept json
// @Produce json
// @Param id path int true "ID бронирования"
// @Param booking body models.Bookings true "Обновленные данные бронирования"
// @Success 200 {object} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 404 {object} map[string]interface{} "Бронирование не найдено"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/{id} [put]
func UpdateBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}
	var booking models.Bookings
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Бронирование не найдено"))
		return
	}

	var input models.Bookings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := db.DB.Model(&booking).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить бронирование"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(booking))
}

// @Summary Удалить бронирование
// @Description Удаляет бронирование по ID
// @Tags Бронирования
// @Param id path int true "ID бронирования"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный ID"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/{id} [delete]
func DeleteBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}

	if err := db.DB.Delete(&models.Bookings{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить бронирование"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Бронирование успешно удалено"))
}

// @Summary Получить бронирования клиента
// @Description Получает список бронирований по ID клиента
// @Tags Бронирования
// @Produce json
// @Param client_id path int true "ID клиента"
// @Success 200 {array} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID клиента"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/client/{client_id} [get]
func GetBookingsByClientHandler(c *gin.Context) {
	clientID, err := strconv.Atoi(c.Param("client_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}

	var bookings []models.Bookings
	if err := db.DB.Where("client_id = ?", clientID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить бронирования"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Бронирование успешно "))
}

// @Summary Получить бронирования услуги
// @Description Получает список бронирований по ID услуги
// @Tags Бронирования
// @Produce json
// @Param service_id path int true "ID услуги"
// @Success 200 {array} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID услуги"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/service/{service_id} [get]
func GetBookingsByServiceHandler(c *gin.Context) {
	serviceID, err := strconv.Atoi(c.Param("service_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid service ID"})
		return
	}

	var bookings []models.Bookings
	if err := db.DB.Where("service_id = ?", serviceID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// @Summary Получить бронирования пользователя
// @Description Получает список бронирований по ID пользователя
// @Tags Бронирования
// @Produce json
// @Param user_id path int true "ID пользователя"
// @Success 200 {array} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID пользователя"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/user/{user_id} [get]
func GetBookingsByUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid user ID"})
		return
	}

	var bookings []models.Bookings
	if err := db.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// @Summary Проверить доступность бронирования
// @Description Проверяет, доступен ли временной слот для пользователя
// @Tags Бронирования
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Param booking_time query string true "Время бронирования"
// @Success 200 {object} map[string]interface{} "Доступность слота"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 409 {object} map[string]interface{} "Слот времени уже занят"
// @Router /bookings/availability [get]
func CheckBookingAvailabilityHandler(c *gin.Context) {
	userID := c.Query("user_id")
	bookingTime := c.Query("booking_time")

	if userID == "" || bookingTime == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("user_id и booking_time обязательны"))
		return
	}

	var count int64
	db.DB.Model(&models.Bookings{}).
		Where("user_id = ? AND booking_time = ?", userID, bookingTime).
		Count(&count)

	if count > 0 {
		c.JSON(http.StatusConflict, utils.SuccessResponse(map[string]bool{"available": false}))
	} else {
		c.JSON(http.StatusOK, utils.SuccessResponse(map[string]bool{"available": true}))
	}
}
