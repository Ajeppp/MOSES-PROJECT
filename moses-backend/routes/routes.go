package routes

import (
	"net/http"

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

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}

	// =========================
	// PROTECTED
	// =========================
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			user, ok := c.Get("user")
			if !ok || user == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
		})
	}
}
