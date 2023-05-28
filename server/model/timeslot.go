package model

import "time"

type Timeslot struct {
	ID      uint64 `gorm:"primaryKey"`
	Date    time.Time
	EventID uint64
	Event   Event
}
