package models

type PlayerRole struct {
	PlayerID uint `gorm:"primaryKey"`
	RoleID   uint `gorm:"primaryKey"`
}
