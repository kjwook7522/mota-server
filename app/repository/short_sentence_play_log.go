package repository

import (
	"gorm.io/gorm"
	"mota-server/app/domain"
)

type ShortSentencePlayLogRepository struct {
	db *gorm.DB
}

func NewShortSentencePlayLogRepository(db *gorm.DB) *ShortSentencePlayLogRepository {
	return &ShortSentencePlayLogRepository{db}
}

func (repo *ShortSentencePlayLogRepository) FindAll(limit, offset int) ([]*domain.ShortSentencePlayLog, error) {
	playLog := make([]*domain.ShortSentencePlayLog, 0)
	err := repo.db.Limit(limit).Offset(offset).Find(&playLog).Error
	return playLog, err
}

func (repo *ShortSentencePlayLogRepository) Insert(model *domain.ShortSentencePlayLog) (uint64, error) {
	err := repo.db.Create(model).Error
	return model.ID, err
}
