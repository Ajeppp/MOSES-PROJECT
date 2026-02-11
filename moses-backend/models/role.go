package models

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;size:50"`
	Name string `gorm:"size:100"`
}
