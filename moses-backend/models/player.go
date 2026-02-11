package models

import "time"

type Player struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`

	MainRoleID uint
	MainRole   Role `gorm:"foreignKey:MainRoleID;references:ID"` // 🔥 FIX

	Roles []PlayerRole `gorm:"foreignKey:PlayerID"` // 🔥 FIX: pivot entity

	Active    bool `gorm:"default:true"`
	CreatedAt time.Time
}
