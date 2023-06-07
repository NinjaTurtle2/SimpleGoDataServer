package handlers

import (
	"myHttpServer/controllers"
	"myHttpServer/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func PostData(c *gin.Context) {
	var data models.Data
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controllers.SaveData(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, data)
}

