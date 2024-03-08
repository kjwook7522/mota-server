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

func (repo *ShortSentenceRecordRepository) FindAll(limit, offset int) ([]*entity.ShortSentenceRecord, error) {
	recordEntities := make([]*entity.ShortSentenceRecord, 0)
	err := repo.db.Limit(limit).Offset(offset).Find(&recordEntities).Error
	return recordEntities, err
}

func (repo *ShortSentenceRecordRepository) Insert(recordEntity *entity.ShortSentenceRecord) (uint64, error) {
	err := repo.db.Create(recordEntity).Error
	return recordEntity.ID, err
}
