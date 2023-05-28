package model

type ParticipantChoice struct {
	ParticipantID uint64 `gorm:"primaryKey"`
	Participant   Participant
	TimeslotID    uint64 `gorm:"primaryKey"`
	Timeslot      Timeslot
	Maybe         bool
}
