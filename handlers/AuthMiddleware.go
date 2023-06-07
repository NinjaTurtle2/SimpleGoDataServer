package handlers

import (
	"myHttpServer/controllers"
	"myHttpServer/utils"
	http "net/http"
	"github.com/gin-gonic/gin"

)

//Auth token middleware for gin
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get token from header
		token := c.Request.Header.Get("Authorization")
		//Check if token is valid
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//Get session from token
		session, err := controllers.GetSession(token)
		if err != nil {
			utils.HandleError(c, err)
			c.Abort()
			return
		}

		user, err := controllers.GetUser(session.Value)

		if err != nil {
			utils.HandleError(c, err)
			c.Abort()
			return
		}

		//Set session in context
		c.Set("user", user)

		//Continue if token is valid
		c.Next()
	}
}