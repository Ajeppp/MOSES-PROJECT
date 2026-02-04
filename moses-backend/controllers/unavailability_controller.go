package controllers

import (
	"net/http"
	"time"

	"moses/database"
	"moses/models"

	"github.com/gin-gonic/gin"
)

type CreateUnavailabilityRequest struct {
	PlayerID    uint   `json:"player_id" binding:"required"`
	ServiceDate string `json:"service_date" binding:"required"`
	Month       int    `json:"month" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Reason      string `json:"reason"`
}

func CreateUnavailability(c *gin.Context) {
	var req CreateUnavailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.ServiceDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format (YYYY-MM-DD)"})
		return
	}

	unav := models.Unavailability{
		PlayerID:    req.PlayerID,
		ServiceDate: date,
		Month:       req.Month,
		Year:        req.Year,
		Reason:      req.Reason,
	}

	database.DB.Create(&unav)
	c.JSON(http.StatusOK, unav)
}

func GetUnavailability(c *gin.Context) {
	var data []models.Unavailability
	database.DB.Preload("Player").Find(&data)
	c.JSON(http.StatusOK, data)
}
