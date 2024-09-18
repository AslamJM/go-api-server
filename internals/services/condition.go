package services

import (
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/repositories"
)

type ConditionService struct {
	repo *repositories.ConditionRepository
}

func NewConditionService(repo *repositories.ConditionRepository) *ConditionService {
	return &ConditionService{repo: repo}
}

func (s *ConditionService) CreateCondition(input *models.CreateConditionInput) error {
	return s.repo.CreateCondition(input)
}
