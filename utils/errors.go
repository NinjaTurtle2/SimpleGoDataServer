package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//AlreadyExistsError is a custom error type function
type AlreadyExistsError struct {
	Message string
}

//Error returns the error message
func (e *AlreadyExistsError) Error() string {
	return e.Message
}

//NotFoundError is a custom error type function
type NotFoundError struct {
	Message string
}

//Error returns the error message
func (e *NotFoundError) Error() string {
	return e.Message
}

//InvalidCredentialsError is a custom error type function
type InvalidCredentialsError struct {
	Message string
}

//Error returns the error message
func (e *InvalidCredentialsError) Error() string {
	return e.Message
}



//Handle error with switch case gin context
func HandleError(c *gin.Context, err error) {
	switch err.(type) {
	case *AlreadyExistsError:
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	case *NotFoundError:
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	case *InvalidCredentialsError:
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	default:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}