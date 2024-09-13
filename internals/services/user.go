package services

import (
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserRepository(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)

}
