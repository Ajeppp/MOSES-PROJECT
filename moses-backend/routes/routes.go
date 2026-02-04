package routes

import (
	"moses/controllers"
	"moses/handlers"
	"moses/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	// players
	api.POST("/players", controllers.CreatePlayer)
	api.GET("/players", controllers.GetPlayers)

	// unavailability
	api.POST("/unavailable", controllers.CreateUnavailability)
	api.GET("/unavailable", controllers.GetUnavailability)

	// scheduler
	api.POST("/generate-schedule", controllers.GenerateSchedule)
	api.GET("/schedules", controllers.GetSchedules)

	// schedule view
	api.GET("/schedule-view", handlers.GetSchedule)

	// debug
	r.GET("/debug", func(c *gin.Context) {
		for _, route := range r.Routes() {
			println(route.Method, route.Path)
		}
	})

	// =========================
	// AUTH
	// =========================
	authHandler := handlers.NewAuthHandler()

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// =========================
	// PROTECTED
	// =========================
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, user)
		})
	}
}
