package handlers

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/services"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookingHandler struct {
	BookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *BookingHandler {
	return &BookingHandler{
		BookingService: bookingService,
	}
}

// @Summary Создать бронирование
// @Description Создает новое бронирование, если слот времени не занят
// @Security BearerAuth
// @Tags Бронирования
// @Accept json
// @Produce json
// @Param booking body models.Bookings true "Данные бронирования"
// @Success 201 {object} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 409 {object} map[string]interface{} "Слот времени уже занят"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings [post]
func (h *BookingHandler) CreateBookingHandler(c *gin.Context) {
	var booking models.Bookings

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.BookingService.CreateBooking(&booking); err != nil {
		if err == services.ErrTimeSlotOccupied {
			c.JSON(http.StatusConflict, utils.ErrorResponse("Временной слот уже занят"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось создать бронирование"))
		}
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(booking))
}

// @Summary Получить бронирование
// @Description Получает бронирование по ID
// @Security BearerAuth
// @Tags Бронирования
// @Produce json
// @Param id path int true "ID бронирования"
// @Success 200 {object} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID"
// @Failure 404 {object} map[string]interface{} "Бронирование не найдено"
// @Router /bookings/{id} [get]
func (h *BookingHandler) GetBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}

	booking, err := h.BookingService.GetBookingByID(id)
	if err != nil {
		if err == services.ErrBookingNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Бронирование не найдено"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при получении бронирования"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(booking))
}

// @Summary Получить все бронирования
// @Security BearerAuth
// @Description Получает список всех бронирований
// @Tags Бронирования
// @Produce json
// @Success 200 {array} models.Bookings
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings [get]
func (h *BookingHandler) GetAllBookingsHandler(c *gin.Context) {
	bookings, err := h.BookingService.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить список бронирований"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(bookings))
}

// @Summary Обновить бронирование
// @Description Обновляет данные бронирования по ID
// @Security BearerAuth
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
func (h *BookingHandler) UpdateBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}

	var input models.Bookings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректные данные: "+err.Error()))
		return
	}

	if err := h.BookingService.UpdateBooking(id, &input); err != nil {
		if err == services.ErrBookingNotFound {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Бронирование не найдено"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось обновить бронирование"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Бронирование успешно обновлено"))
}

// @Summary Удалить бронирование
// @Security BearerAuth
// @Description Удаляет бронирование по ID
// @Tags Бронирования
// @Param id path int true "ID бронирования"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Некорректный ID"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/{id} [delete]
func (h *BookingHandler) DeleteBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID бронирования"))
		return
	}

	if err := h.BookingService.DeleteBooking(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось удалить бронирование"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Бронирование успешно удалено"))
}

// @Summary Проверить доступность бронирования
// @Security BearerAuth
// @Description Проверяет, доступен ли временной слот для пользователя
// @Tags Бронирования
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Param booking_time query string true "Время бронирования"
// @Success 200 {object} map[string]interface{} "Доступность слота"
// @Failure 400 {object} map[string]interface{} "Некорректный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/availability [get]
func (h *BookingHandler) CheckBookingAvailabilityHandler(c *gin.Context) {
	userIDStr := c.Query("user_id")
	bookingTime := c.Query("booking_time")

	if userIDStr == "" || bookingTime == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("user_id и booking_time обязательны"))
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный user_id"))
		return
	}

	available, err := h.BookingService.CheckAvailability(userID, bookingTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Ошибка при проверке доступности"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(map[string]bool{"available": available}))
}

// @Summary Получить бронирования клиента
// @Security BearerAuth
// @Description Получает список бронирований по ID клиента
// @Tags Бронирования
// @Produce json
// @Param client_id path int true "ID клиента"
// @Success 200 {array} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID клиента"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/client/{client_id} [get]
func (h *BookingHandler) GetBookingsByClientHandler(c *gin.Context) {
	clientID, err := strconv.Atoi(c.Param("client_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID клиента"))
		return
	}

	bookings, err := h.BookingService.GetBookingsByClientID(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить бронирования"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(bookings))
}

// @Summary Получить бронирования услуги
// @Security BearerAuth
// @Description Получает список бронирований по ID услуги
// @Tags Бронирования
// @Produce json
// @Param service_id path int true "ID услуги"
// @Success 200 {array} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID услуги"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/service/{service_id} [get]
func (h *BookingHandler) GetBookingsByServiceHandler(c *gin.Context) {
	serviceID, err := strconv.Atoi(c.Param("service_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID услуги"))
		return
	}

	bookings, err := h.BookingService.GetBookingsByServiceID(serviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить бронирования"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(bookings))
}

// @Summary Получить бронирования пользователя
// @Security BearerAuth
// @Description Получает список бронирований по ID пользователя
// @Tags Бронирования
// @Produce json
// @Param user_id path int true "ID пользователя"
// @Success 200 {array} models.Bookings
// @Failure 400 {object} map[string]interface{} "Некорректный ID пользователя"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /bookings/user/{user_id} [get]
func (h *BookingHandler) GetBookingsByUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Некорректный ID пользователя"))
		return
	}

	bookings, err := h.BookingService.GetBookingsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Не удалось получить бронирования"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(bookings))
}
