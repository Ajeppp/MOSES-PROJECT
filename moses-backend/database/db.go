package database

import (
	"log"
	"moses/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=jef dbname=moses port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = db.AutoMigrate(
		&models.Player{},
		&models.Unavailability{},
		&models.ServiceSchedule{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	DB = db
	log.Println("Database connected")
}
