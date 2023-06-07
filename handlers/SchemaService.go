package handlers

import (
	"fmt"
	"myHttpServer/controllers"
	"myHttpServer/models"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// getSchemas responds with the list of all schemas as JSON.
func GetSchemas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, controllers.GetSchemas())
}

// GetSchemaByID locates the schema whose ID value matches the id
func GetSchemaByID(c *gin.Context) {
    id := c.Param("id")
	schema := controllers.GetSchemaByID(id)
	if schema != nil {
		c.IndentedJSON(http.StatusOK, schema)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostSchemas(c *gin.Context) {
    var newSchema models.Schema

    // Call BindJSON to bind the received JSON to
    // newSchema.
    if err := c.BindJSON(&newSchema); err != nil {
		fmt.Println(err)
	    return
    }

    // Add the new album to the slice.
    controllers.SaveSchema(&newSchema)
    c.IndentedJSON(http.StatusCreated, newSchema)
}