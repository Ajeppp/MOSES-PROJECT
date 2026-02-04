package controllers

import (
	"net/http"

	"moses/database"
	"moses/models"
	"moses/scheduler"

	"github.com/gin-gonic/gin"
)

type GenerateRequest struct {
	Month int `json:"month" binding:"required"`
	Year  int `json:"year" binding:"required"`
}

func GenerateSchedule(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := scheduler.GenerateSchedule(req.Month, req.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "schedule generated"})
}

func GetSchedules(c *gin.Context) {
	var schedules []models.ServiceSchedule
	database.DB.Preload("Player").Find(&schedules)
	c.JSON(http.StatusOK, schedules)
}
