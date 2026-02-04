package models

import "time"

type ServiceSchedule struct {
	ID          uint      `gorm:"primaryKey"`
	ServiceDate time.Time `gorm:"not null"`
	Month       int       `gorm:"not null"`
	Year        int       `gorm:"not null"`
	Role        string    `gorm:"size:50;not null"`
	PlayerID    uint
	CreatedAt   time.Time

	Player Player `gorm:"foreignKey:PlayerID"`
}
