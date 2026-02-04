package main

import (
	"moses/database"
	"moses/routes"
	"moses/scheduler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ✅ CORS CONFIG
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	database.Connect()
	routes.RegisterRoutes(r)

	r.Run(":8080")

	scheduler.GenerateSchedule(2, 2026)
}
