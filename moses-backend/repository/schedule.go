package repository

import (
	"moses/database"
	"moses/models"
)

func GetSchedules(month int, year int) ([]models.ServiceSchedule, error) {
	db := database.DB

	var schedules []models.ServiceSchedule
	err := db.Preload("Player").
		Where("month = ? AND year = ?", month, year).
		Order("service_date asc").
		Find(&schedules).Error

	return schedules, err
}
