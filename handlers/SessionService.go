package handlers

import (
	"myHttpServer/controllers"
	"myHttpServer/utils"
	http "net/http"
	"github.com/gin-gonic/gin"

)

//Get session by key gin
func GetSession(c *gin.Context) {
	key := c.Param("key")
	session, err := controllers.GetSession(key)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, session)
}

//Create session gin
func CreateSession(c *gin.Context) {
	username := c.Param("username")
	user, err := controllers.GetUserByUsername(username)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	session, err := controllers.CreateSession(user)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, session)
}
