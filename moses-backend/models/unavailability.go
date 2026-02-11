package models

import "time"

type Unavailability struct {
	ID uint `gorm:"primaryKey"`

	PlayerID uint
	Player   Player `gorm:"foreignKey:PlayerID"`

	ServiceDate time.Time
	Month       int
	Year        int
	Reason      string
	CreatedAt   time.Time
}
