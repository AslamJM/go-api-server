package services

import "kstrategy-backend/internals/repositories"

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(r *repositories.UserRepository) *AuthService {
	return &AuthService{repo: r}
}
