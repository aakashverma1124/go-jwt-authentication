package routes

import (
	controller "github.com/aakashverma1124/go-jwt-authentication/controllers"
	"github.com/gin-gonic/gin"
	"github.com/aakashverma1124/go-jwt-authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
