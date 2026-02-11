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
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	date, err := time.Parse("2006-01-02", req.ServiceDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid date format",
			"error":   "Use YYYY-MM-DD",
		})
		return
	}

	unav := models.Unavailability{
		PlayerID:    req.PlayerID,
		ServiceDate: date,
		Month:       req.Month,
		Year:        req.Year,
		Reason:      req.Reason,
	}

	if err := database.DB.Create(&unav).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create unavailability",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    unav,
	})
}

func GetUnavailability(c *gin.Context) {
	var data []models.Unavailability

	if err := database.DB.
		Preload("Player").
		Preload("Player.MainRole").
		Preload("Player.Roles").
		Preload("Player.Roles.Role").
		Find(&data).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed fetch unavailability",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
