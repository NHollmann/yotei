package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"not null;size:100"`
	Username     string `gorm:"not null;size:100"`
	passwordHash string `gorm:"not null;size:60"`
	IsAdmin      bool   `gorm:"not null"`
}
