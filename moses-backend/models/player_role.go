package models

type PlayerRole struct {
	ID uint `gorm:"primaryKey"`

	PlayerID uint
	Player   Player `gorm:"foreignKey:PlayerID"`

	RoleID uint
	Role   Role `gorm:"foreignKey:RoleID"`

	Type string `gorm:"size:20"` // MAIN / ADDITIONAL
}
