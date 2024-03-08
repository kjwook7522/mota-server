package repository

import (
	"gorm.io/gorm"
	"mota-server/app/entity"
)

type ShortSentenceRepository struct {
	db *gorm.DB
}

func NewShortSentenceRepository(db *gorm.DB) *ShortSentenceRepository {
	return &ShortSentenceRepository{db: db}
}

func (repo *ShortSentenceRepository) FindAll(limit, offset int) ([]*entity.ShortSentenceEntity, error) {
	logEntities := make([]*entity.ShortSentenceEntity, 0)
	err := repo.db.Limit(limit).Offset(offset).Find(&logEntities).Error
	return logEntities, err
}
