package services

import (
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/repositories"
	"kstrategy-backend/internals/utils"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(r *repositories.UserRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Login(input models.LoginInput) (*models.LoginResponse, error) {
	user, err := s.repo.GetUserByUsername(input.Username)

	if err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateAccessToken(int(user.ID))

	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		AccessToken: accessToken,
		User:        *user,
	}, nil
}
