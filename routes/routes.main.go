package routes

import (
	"os"

	"jual-beli-motor/controllers"
	"jual-beli-motor/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.GET("/ping", controllers.HealthCheck)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/register", controllers.CreateUserNonAdmin)

	bikeType := r.Group("/bike-type")
	{
		bikeType.POST("/", controllers.CreateBikeType)
		bikeType.GET("/", controllers.GetAllBikeType)
		bikeType.GET("/:id", controllers.GetBikeTypeById)
		bikeType.PUT("/:id", controllers.UpdateBikeType)
		bikeType.DELETE("/:id", controllers.DeleteBikeTypeById)
	}
	// Admin -> (non-aktif user, create bike type, update bike type, delete bike type)
	// User (create bike, update, delete)
	// Awam (get all bike, get detail bike)
	bike := r.Group("/bike")
	{
		bike.POST("/", middleware.Authentication("user"), controllers.CreateBike)
		bike.GET("/", middleware.Authentication("user"), controllers.GetAllBike)
	}

	r.Run(os.Getenv("PORT"))
}
