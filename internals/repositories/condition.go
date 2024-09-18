package repositories

import (
	"kstrategy-backend/internals/models"

	"gorm.io/gorm"
)

type ConditionRepository struct {
	db *gorm.DB
}

func NewConditionRepository(db *gorm.DB) *ConditionRepository {
	return &ConditionRepository{db: db}
}

func (r *ConditionRepository) CreateCondition(input *models.CreateConditionInput) error {
	return r.db.Create(input).Error
}
