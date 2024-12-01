package services

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func CreateServiceHandler(c *gin.Context) {
	var service models.Services
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed create service"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func GetAllServicesHandler(c *gin.Context) {
	var services []models.Services
	if err := db.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get services"})
		return
	}
	c.JSON(http.StatusOK, services)
}

func GetServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service id"})
		return
	}

	var service models.Services
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

func UpdateServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service id"})
		return
	}

	var service models.Services
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "service not found"})
		return
	}

	var input models.Services
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&service).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed update service"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func DeleteServiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service id"})
		return
	}

	if err := db.DB.Delete(&models.Services{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed delete service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "service deleted"})
}
