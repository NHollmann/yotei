package model

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"gorm.io/gorm"
)

const AccessKeyLength = 10

type Event struct {
	gorm.Model
	Name         string        `gorm:"not null;size:200"`
	CreatedByID  uint          `gorm:"not null"`
	CreatedBy    User          ``
	AccessKey    string        `gorm:"unique;not null;size:10"`
	Timeslots    []Timeslot    ``
	Participants []Participant ``
}

func EventGetAll(db *gorm.DB, createdById uint) []Event {
	var events []Event
	var result *gorm.DB
	if createdById != 0 {
		result = db.Where("created_by_id = ?", createdById).Find(&events)
	} else {
		result = db.Find(&events)
	}
	if result.Error != nil {
		log.Panic("error on loading all events", result.Error)
	}

	return events
}

func EventGetOne(db *gorm.DB, accessKey string) (Event, error) {
	var event Event
	result := db.Where("access_key = ?", accessKey).First(&event)
	if result.Error != nil {
		log.Print("db error on EventGetOne", result.Error)
		return Event{}, fmt.Errorf("event not found: %s", accessKey)
	}

	return event, nil
}

func EventCreate(db *gorm.DB, name string, createdByID uint) (string, error) {

	var user User
	result := db.First(&user, createdByID)
	if result.Error != nil {
		log.Print("db error on EventCreate", result.Error)
		return "", fmt.Errorf("user not found: %d", createdByID)
	}
	event := Event{
		Name:        strings.TrimSpace(name),
		CreatedByID: createdByID,
		AccessKey:   generateAccessKey(),
	}

	result = db.Create(&event)
	if result.Error != nil {
		log.Print("db error on EventCreate", result.Error)
		return "", fmt.Errorf("event not created")
	}

	return event.AccessKey, nil
}

func EventDelete(db *gorm.DB, accessKey string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		var event Event
		result := tx.Where("access_key = ?", accessKey).First(&event)
		if result.Error != nil {
			log.Print("db error on EventDelete", result.Error)
			return fmt.Errorf("event not found: %s", accessKey)
		}

		result = tx.Delete(&event)
		if result.Error != nil {
			log.Print("db error on EventDelete", result.Error)
			return fmt.Errorf("event not deleted")
		}
		return nil
	})
}

func generateAccessKey() string {
	var letters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	runeList := make([]rune, AccessKeyLength)

	for i := range runeList {
		runeList[i] = letters[rand.Intn(len(letters))]
	}

	return string(runeList)
}
