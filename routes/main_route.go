package routes

import (
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/controllers"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	// Collection of routes for user profile
	route.POST("/users/register", controllers.RegisterUserAccount)
	route.GET("/users/login", controllers.LoginUserAccount)
	route.POST("/users/logout", controllers.LogoutUserAccount)

	auth_route := route.Group("/")

	auth_route.Use(middlewares.JwtCheck())
	{
		// Collection of routes for user profile (specifically to put and delete)
		auth_route.PUT("/users/:id", controllers.UpdateUserAccount)
		auth_route.DELETE("/users/:id", controllers.DeleteUserAccount)

		// Collection of routes for user photo
		auth_route.POST("/photos", controllers.CreateUserPhoto)
		auth_route.GET("/photos", controllers.GetUserPhoto)
		auth_route.PUT("/photos/:id", middlewares.PhotoAuthorization(), controllers.UpdateUserPhoto)
		auth_route.DELETE("/photos/:id", middlewares.PhotoAuthorization(), controllers.DeleteUserPhoto)
	}

	return route
}
