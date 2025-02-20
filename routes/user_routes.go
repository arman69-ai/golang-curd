package routes

import (
	"github.com/arman69-ai/golang-crud/controllers"
	"github.com/arman69-ai/golang-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// Lindungi semua route user
	protected := router.Group("/users")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("", controllers.GetUsers)
		protected.GET("/:id", controllers.GetUser)
		protected.PUT("/:id", controllers.UpdateUser)
		protected.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}
}
