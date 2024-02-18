package domain

import "time"

type ShortSentencePlayLog struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	IpAddress  string    `json:"ip_address" gorm:"size:15"`
	UserName   string    `json:"user_name" gorm:"size:50"`
	SentenceID uint64    `json:"sentence_id"`
	Typing     string    `json:"typing"`
	Speed      int       `json:"speed"`
	CreatedAt  time.Time `json:"created_at"`
}
