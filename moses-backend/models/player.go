package models

import "time"

type Player struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Role      string `gorm:"size:50;not null"`
	Active    bool   `gorm:"default:true"`
	CreatedAt time.Time
}
