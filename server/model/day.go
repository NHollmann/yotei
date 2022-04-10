package model

import "time"

type Day struct {
	ID      uint64 `gorm:"primaryKey"`
	Date    time.Time
	EventID uint64
	Event   Event
}
