package repository

import (
	"gorm.io/gorm"
	"mota-server/app/entity"
)

type ShortSentenceRecordRepository struct {
	db *gorm.DB
}

func NewShortSentenceRecordRepository(db *gorm.DB) *ShortSentenceRecordRepository {
	return &ShortSentenceRecordRepository{db}
}

func (repo *ShortSentenceRecordRepository) FindAll(limit, offset int) ([]*entity.ShortSentenceRecordEntity, error) {
	recordEntities := make([]*entity.ShortSentenceRecordEntity, 0)
	err := repo.db.Limit(limit).Offset(offset).Find(&recordEntities).Error
	return recordEntities, err
}

func (repo *ShortSentenceRecordRepository) Insert(recordEntity *entity.ShortSentenceRecordEntity) (uint64, error) {
	err := repo.db.Create(recordEntity).Error
	return recordEntity.ID, err
}
