package breaks

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func CreateBreakHandler(c *gin.Context) {
	var breaks models.Breaks
	if err := c.ShouldBindJSON(&breaks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&breaks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed create break"})
		return
	}

	c.JSON(http.StatusOK, breaks)

}

func GetAllBreaksHandler(c *gin.Context) {
	var breaks []models.Breaks
	if err := db.DB.Find(&breaks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get all breaks"})
		return
	}
	c.JSON(http.StatusOK, breaks)
}

func GetBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var breaks models.Breaks
	if err := db.DB.First(&breaks, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
		return
	}

	c.JSON(http.StatusOK, breaks)
}

func UpdateBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var breaks models.Breaks
	if err := db.DB.First(&breaks, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "break not found"})
		return
	}

	var input models.Breaks
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&breaks).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to update break"})
		return
	}

	c.JSON(http.StatusOK, breaks)

}

func DeleteBreakHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := db.DB.Delete(&models.Breaks{}).Delete(&models.Breaks{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "break deleted successfully"})
}
