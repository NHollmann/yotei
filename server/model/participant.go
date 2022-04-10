package model

type Participant struct {
	ID      uint64 `gorm:"primaryKey"`
	Name    string
	UserID  uint
	User    *User
	EventID uint64
	Event   Event
}
