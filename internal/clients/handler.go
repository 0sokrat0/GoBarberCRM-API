package clients

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func CreateClientHandler(c *gin.Context) {
	var client models.Clients

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create client"})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func GetAllClientsHandler(c *gin.Context) {
	var clients []models.Clients

	if err := db.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func GetClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get client"})
		return
	}
	var client models.Clients
	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

func UpdateClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	var client models.Clients
	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}

	var input models.Clients
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to update client"})
		return
	}
	c.JSON(http.StatusOK, client)
}

func DeleteClientHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	if err := db.DB.Delete(&models.Clients{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete client"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "client deleted successfully"})
}
