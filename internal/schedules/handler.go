package schedules

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func CreateScheduleHandler(c *gin.Context) {
	var schedule models.Schedules
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create schedule"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func GetAllSchedulesHandler(c *gin.Context) {
	var schedules []models.Schedules
	if err := db.DB.Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch schedules"})
		return
	}
	c.JSON(http.StatusOK, schedules)
}

func GetScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var schedule models.Schedules
	if err := db.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch schedule"})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

func UpdateScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var schedule models.Schedules
	if err := db.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch schedule"})
		return
	}

	var input models.Schedules
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&schedule).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update schedule"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func DeleteScheduleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := db.DB.Delete(&models.Schedules{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": " schedule deleted successfully"})
}
