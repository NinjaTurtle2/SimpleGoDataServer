package handlers

import (
	"myHttpServer/controllers"
	"myHttpServer/models"
	"myHttpServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exists := c.Get(utils.CONTEXT_CURRENT_USER)
	username := user.(models.User).Username
	if !exists || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found in context"})
		return
	}
	err = controllers.SaveTask(&username, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func TaskMongoWebhook(c *gin.Context) {
	taskId := c.Param(utils.TASK_ID)
	if taskId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := controllers.SyncTaskToSheet(&taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "SUCCESS")
}
