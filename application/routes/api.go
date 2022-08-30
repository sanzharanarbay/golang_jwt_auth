package routes

import (
	"github.com/gin-gonic/gin"
	usersController "jwt_auth_golang/application/controllers/users"
	authController "jwt_auth_golang/application/controllers/auth"
	"jwt_auth_golang/application/middleware"
)

func ApiRoutes(prefix string, router *gin.Engine) {

	// general routes
		router.POST("/login", authController.Login)
		router.POST("/logout", authController.Logout)
		router.POST("/refresh", authController.Refresh)
		router.POST("/register", usersController.CreateUser)

		// dashboard routes

	apiGroup := router.Group(prefix)
	{
		dashboard := apiGroup.Group("/dashboard").Use(middleware.TokenAuthMiddleware())
		{
			dashboard.GET("/users", usersController.GetUsers)
			dashboard.GET("/users/:id", usersController.GetUser)
			dashboard.POST("/users", usersController.CreateUser)
			dashboard.PUT("/users/:id", usersController.UpdateUser)
			dashboard.DELETE("/users/:id", usersController.DeleteUser)
		}
	}


}
