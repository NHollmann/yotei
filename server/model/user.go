package model

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const passwordHashCost = 10

type User struct {
	gorm.Model
	Name                   string `gorm:"not null;size:100"`
	Username               string `gorm:"unique;not null;size:100"`
	passwordHash           string `gorm:"not null;size:60"`
	IsAdmin                bool   `gorm:"not null"`
	PasswordChangeRequired bool   `gorm:"not null;default:0"`
}

func (u *User) BeforeDelete(db *gorm.DB) (err error) {
	if u.IsAdmin {
		return errors.New("admin user not allowed to delete")
	}
	return
}

func UserGetAll(db *gorm.DB) []User {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Panic("error on loading all users", result.Error)
	}

	return users
}

func UserGetOne(db *gorm.DB, id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		log.Print("db error on UserGetOne", result.Error)
		return User{}, fmt.Errorf("user not found: %d", id)
	}

	return user, nil
}

func UserCreate(db *gorm.DB, name, username, password string, isAdmin, passwordChangeRequired bool) (uint, error) {

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	if err != nil {
		log.Print("password hash error on UserCreate", err)
		return 0, fmt.Errorf("user not created")
	}

	user := User{
		Name:                   strings.TrimSpace(name),
		Username:               strings.TrimSpace(username),
		passwordHash:           string(hashBytes),
		IsAdmin:                isAdmin,
		PasswordChangeRequired: passwordChangeRequired,
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Print("db error on UserCreate", result.Error)
		return 0, fmt.Errorf("user not created")
	}

	return user.ID, nil
}

func UserUpdate(db *gorm.DB, id uint, name, username, password string, isAdmin, passwordChangeRequired bool) error {

	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		log.Print("db error on UserUpdate", result.Error)
		return fmt.Errorf("user not found: %d", id)
	}

	if password != "" {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
		if err != nil {
			log.Print("password hash error on UserUpdate", err)
			return fmt.Errorf("user not updated")
		}
		user.passwordHash = string(hashBytes)
	}

	user.Name = name
	user.Username = username
	user.IsAdmin = isAdmin
	user.PasswordChangeRequired = passwordChangeRequired

	result = db.Save(&user)
	if result.Error != nil {
		log.Print("db error on UserUpdate", result.Error)
		return fmt.Errorf("user not updated")
	}

	return nil
}

func UserDelete(db *gorm.DB, id uint) error {
	return db.Transaction(func(tx *gorm.DB) error {

		var user User
		result := tx.First(&user, id)
		if result.Error != nil {
			log.Print("db error on UserDelete", result.Error)
			return fmt.Errorf("user not found: %d", id)
		}

		result = tx.Delete(&user)
		if result.Error != nil {
			log.Print("db error on UserDelete", result.Error)
			return fmt.Errorf("user not deleted")
		}
		return nil
	})
}

func UserCheckLogin(db *gorm.DB, username, password string) bool {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Print("db error on UserCheckLogin", result.Error)
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.passwordHash), []byte(password))
	return err == nil
}
