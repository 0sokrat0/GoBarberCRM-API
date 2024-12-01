package bookings

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func CreateBookingHandler(c *gin.Context) {
	var booking models.Bookings

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func GetBookingHandler(c *gin.Context) {
	var booking models.Bookings
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "booking not found"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func GetAllBookingsHandler(c *gin.Context) {
	var bookings []models.Bookings
	if err := db.DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func UpdateBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}
	var booking models.Bookings
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking is not found"})
		return
	}

	var input models.Bookings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&booking).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update booking"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func DeleteBookingHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	if err := db.DB.Delete(&models.Bookings{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking deleted successfully"})
}
