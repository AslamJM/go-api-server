package services

import (
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewRepository(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.repo.GetAllUsers()
}
