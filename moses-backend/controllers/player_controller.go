package controllers

import (
	"net/http"

	"moses/database"
	"moses/models"

	"github.com/gin-gonic/gin"
)

type CreatePlayerReq struct {
	Name            string   `json:"name" binding:"required"`
	MainRole        string   `json:"main_role" binding:"required"`
	AdditionalRoles []string `json:"additional_roles"`
}

func CreatePlayer(c *gin.Context) {
	var req CreatePlayerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.AdditionalRoles) > 2 {
		c.JSON(400, gin.H{"error": "max 2 additional roles allowed"})
		return
	}

	db := database.DB

	// 🔹 get main role
	var mainRole models.Role
	if err := db.Where("code = ?", req.MainRole).First(&mainRole).Error; err != nil {
		c.JSON(400, gin.H{"error": "main role not found"})
		return
	}

	// 🔹 create player
	player := models.Player{
		Name:       req.Name,
		MainRoleID: mainRole.ID,
		Active:     true,
	}

	if err := db.Create(&player).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 🔹 additional roles
	var additionalRoles []models.Role
	for _, roleCode := range req.AdditionalRoles {
		if roleCode == req.MainRole {
			continue // prevent duplicate main role
		}

		var role models.Role
		if err := db.Where("code = ?", roleCode).First(&role).Error; err == nil {
			additionalRoles = append(additionalRoles, role)
		}
	}

	if len(additionalRoles) > 0 {
		db.Model(&player).Association("Roles").Append(&additionalRoles)
	}

	c.JSON(201, gin.H{
		"message": "player created",
		"player": gin.H{
			"id":   player.ID,
			"name": player.Name,
			"roles": gin.H{
				"main":       mainRole.Code,
				"additional": req.AdditionalRoles,
			},
		},
	})
}

func GetPlayers(c *gin.Context) {
	db := database.DB

	var players []models.Player
	db.Preload("MainRole").Preload("Roles").Find(&players)

	var result []gin.H

	for _, p := range players {
		additional := []string{}
		for _, r := range p.Roles {
			additional = append(additional, r.Role.Code)
		}

		result = append(result, gin.H{
			"id":   p.ID,
			"name": p.Name,
			"roles": gin.H{
				"main":       p.MainRole.Code,
				"additional": additional,
			},
		})
	}

	c.JSON(200, result)
}
