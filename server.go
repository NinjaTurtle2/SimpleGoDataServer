package main

import (
	"fmt"
	"myHttpServer/handlers"

	gin "github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	//Public Routes
	public := r.Group("/api")
	public.POST("/login", handlers.LoginUser)
	public.POST("/users/:special", handlers.CreateUser)
	

	//Private Routes
	private := r.Group("/api")
	private.Use(handlers.SessionMiddleware())

	private.GET("/schemas", handlers.GetSchemas)
	private.GET("/schemas/:id", handlers.GetSchemaByID)
	private.POST("/schemas", handlers.PostSchemas)

	private.GET("/users", handlers.GetUsers)
	private.GET("/users/:username", handlers.GetUserByUsername)
	private.DELETE("/users/:id", handlers.DeleteUser)

	private.POST("/data", handlers.PostData)

	private.POST("/task", handlers.PostTask)
	
	//router.GET("/sessions/:key", handlers.GetSession)
	//router.POST("/sessions/:username", handlers.CreateSession)

	err:= r.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
