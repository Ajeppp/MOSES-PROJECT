package controllers

import (
	"net/http"

	"moses/database"
	"moses/models"

	"github.com/gin-gonic/gin"
)

type CreatePlayerRequest struct {
	Name string `json:"name" binding:"required"`
	Role string `json:"role" binding:"required"`
}

func CreatePlayer(c *gin.Context) {
	var req CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player := models.Player{
		Name:   req.Name,
		Role:   req.Role,
		Active: true,
	}

	database.DB.Create(&player)
	c.JSON(http.StatusOK, player)
}

func GetPlayers(c *gin.Context) {
	var players []models.Player
	database.DB.Find(&players)
	c.JSON(http.StatusOK, players)
}
