package repository

import (
	"gorm.io/gorm"
	"mota-server/app/domain"
)

type ShortSentenceRepository struct {
	db *gorm.DB
}

func NewShortSentenceRepository(db *gorm.DB) *ShortSentenceRepository {
	return &ShortSentenceRepository{db: db}
}

func (repo *ShortSentenceRepository) FindAll(limit, offset int) []*domain.ShortSentence {
	shortSentence := make([]*domain.ShortSentence, 0)
	repo.db.Limit(limit).Offset(offset).Find(&shortSentence)
	return shortSentence
}
