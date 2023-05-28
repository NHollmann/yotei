package model

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string        `gorm:"not null;size:200"`
	CreatedByID  uint          `gorm:"not null"`
	CreatedBy    User          ``
	AccessKey    string        `gorm:"unique;not null;size:10"`
	Timeslots    []Timeslot    ``
	Participants []Participant ``
}
