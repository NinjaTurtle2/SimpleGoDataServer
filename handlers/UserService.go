package handlers

import (
	"myHttpServer/controllers"
	"myHttpServer/models"
	"myHttpServer/utils"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	specialToken := c.Param("special")
	if specialToken != "special" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var newUser models.User

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user, error := controllers.CreateUser(&newUser)
	if error != nil {
		utils.HandleError(c, error)
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, error := controllers.GetUser(id)
	if error != nil {
		utils.HandleError(c, error)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, controllers.GetUsers())
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user,error := controllers.GetUserByUsername(username)
	if user == nil {
		utils.HandleError(c, error)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

//Delete user gin
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	controllers.DeleteUser(id)
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	session, err := controllers.LoginUser(&user)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, session)
}

func UserMongoWebhook(c *gin.Context) {
	userId := c.Param(utils.USER_ID)
	if userId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := controllers.SyncUserToSheet(&userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "SUCCESS")
}


