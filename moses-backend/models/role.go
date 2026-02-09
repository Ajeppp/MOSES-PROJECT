package models

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"size:50;unique;not null"` // WL, SGR, etc
	Name string `gorm:"size:100"`
}
