package domain

import "time"

type ShortSentencePlayLog struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	IpAddress  string `gorm:"size:15"`
	UserName   string `gorm:"size:50"`
	SentenceID uint64
	Typing     string
	Speed      int
	CreatedAt  time.Time
}
