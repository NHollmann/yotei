package model

type ParticipantDay struct {
	ParticipantID uint64 `gorm:"primaryKey"`
	Participant   Participant
	DayID         uint64 `gorm:"primaryKey"`
	Day           Day
	Maybe         bool
}
