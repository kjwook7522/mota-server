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

func (repo *ShortSentenceRepository) FindAll(limit, offset int) ([]*domain.ShortSentence, error) {
	shortSentence := make([]*domain.ShortSentence, 0)
	err := repo.db.Limit(limit).Offset(offset).Find(&shortSentence).Error
	return shortSentence, err
}
