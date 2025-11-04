package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mikehquan19/connect/controllers"
)

func RegisterUserRoutes(r *gin.Engine) {
	user := r.Group("/api/users")

	user.GET("/", controllers.GetUsers)
	user.POST("/", controllers.CreateUser)
}
