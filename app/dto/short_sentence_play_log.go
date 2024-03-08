package dto

import "time"

type ShortSentencePlayLogDto struct {
	ID         uint64    `json:"id"`
	IpAddress  string    `json:"ip_address"`
	UserName   string    `json:"user_name"`
	SentenceID uint64    `json:"sentence_id"`
	Typing     string    `json:"typing"`
	Speed      int       `json:"speed"`
	CreatedAt  time.Time `json:"created_at"`
}
