package main

import (
	"moses/database"
	"moses/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect() // connect DB + automigrate
	routes.RegisterRoutes(r)

	r.Run(":8080") // start server
}
