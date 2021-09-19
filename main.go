package main

import (
	"os"
	"github.com/gin-gonic/gin"
	routes "github.com/aakashverma1124/go-jwt-authentication/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.GET("/api-1", func(context *gin.Context) {
		context.JSON(200, gin.H{"success" : "Access granted for api-1"})
	})
	
	router.GET("/api-2", func(context *gin.Context) {
		context.JSON(200, gin.H{"success" : "Access granted for api-2"})
	})

	router.Run(":" + port)
}