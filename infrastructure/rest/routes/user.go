// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	userController "github.com/phoenixTW/microservices-go/infrastructure/rest/controllers/user"
	"github.com/phoenixTW/microservices-go/infrastructure/rest/middlewares"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerAuth := router.Group("/user")
	routerAuth.Use(middlewares.AuthJWTMiddleware())
	{
		routerAuth.POST("/", controller.NewUser)
		routerAuth.GET("/:id", controller.GetUsersByID)
		routerAuth.GET("/", controller.GetAllUsers)
		routerAuth.PUT("/:id", controller.UpdateUser)
		routerAuth.DELETE("/:id", controller.DeleteUser)
	}
}
