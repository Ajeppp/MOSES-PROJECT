package models

import "time"

type Unavailability struct {
	ID          uint      `gorm:"primaryKey"`
	PlayerID    uint      `gorm:"not null"`
	ServiceDate time.Time `gorm:"not null"`
	Month       int       `gorm:"not null"`
	Year        int       `gorm:"not null"`
	Reason      string    `gorm:"size:255"`
	CreatedAt   time.Time

	Player Player `gorm:"foreignKey:PlayerID"`
}
