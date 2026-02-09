package models

import "time"

type Player struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`

	MainRoleID uint
	MainRole   Role `gorm:"foreignKey:MainRoleID"`

	Roles []Role `gorm:"many2many:player_roles;"` // additional roles (max 2)

	Active    bool `gorm:"default:true"`
	CreatedAt time.Time
}
