package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                   string `gorm:"not null;size:100"`
	Username               string `gorm:"unique;not null;size:100"`
	passwordHash           string `gorm:"not null;size:60"`
	IsAdmin                bool   `gorm:"not null"`
	PasswordChangeRequired bool   `gorm:"not null;default:0"`
}

func UserGetAll(db *gorm.DB) []User {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Panic("Error on loading all users", result.Error)
	}

	return users
}

func UserGetOne(db *gorm.DB, id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		log.Print("DB-Error on GetUser", result.Error)
		return User{}, fmt.Errorf("user not found: %d", id)
	}

	return user, nil
}
