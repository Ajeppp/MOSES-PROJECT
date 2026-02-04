package routes

import (
	"moses/controllers"
	"moses/handlers"

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

	r.GET("/debug", func(c *gin.Context) {
		for _, route := range r.Routes() {
			println(route.Method, route.Path)
		}
	})

}
