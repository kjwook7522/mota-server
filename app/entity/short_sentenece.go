package entity

import "time"

type ShortSentence struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Sentence  string `gorm:"size:500"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}
