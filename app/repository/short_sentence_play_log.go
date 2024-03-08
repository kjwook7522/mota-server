package repository

import (
	"gorm.io/gorm"
	"mota-server/app/entity"
)

type ShortSentencePlayLogRepository struct {
	db *gorm.DB
}

func NewShortSentencePlayLogRepository(db *gorm.DB) *ShortSentencePlayLogRepository {
	return &ShortSentencePlayLogRepository{db}
}

func (repo *ShortSentencePlayLogRepository) FindAll(limit, offset int) ([]*entity.ShortSentencePlayLog, error) {
	logEntities := make([]*entity.ShortSentencePlayLog, 0)
	err := repo.db.Limit(limit).Offset(offset).Find(&logEntities).Error
	return logEntities, err
}

func (repo *ShortSentencePlayLogRepository) Insert(logEntity *entity.ShortSentencePlayLog) (uint64, error) {
	err := repo.db.Create(logEntity).Error
	return logEntity.ID, err
}
