package handlers

import (
	"moses/database"
	"moses/models"
	service "moses/services"

	"github.com/gin-gonic/gin"
)

func GetSchedule(c *gin.Context) {

	var schedules []models.ServiceSchedule

	database.DB.
		Preload("Player").
		Find(&schedules)

	view := service.BuildScheduleView(schedules)

	c.JSON(200, view)
}
