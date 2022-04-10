package model

import "time"

type Event struct {
	ID           uint64        `gorm:"primaryKey"`
	CreatedAt    time.Time     ``
	UpdatedAt    time.Time     ``
	Name         string        `gorm:"not null;size:200"`
	UserID       uint          `gorm:"not null"`
	User         User          ``
	DeletionDate time.Time     `gorm:"not null"`
	AccessKey    string        `gorm:"not null;size:36"`
	Days         []Day         ``
	Participants []Participant ``
}
