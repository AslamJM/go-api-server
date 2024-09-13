package services

import (
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(input *models.UserCreateInput) error {

	hash, err := hashPassword(input.Password)
	if err != nil {
		return err
	}

	user :=
		models.User{
			Username: input.Username,
			Email:    input.Email,
			Password: hash,
		}

	return s.repo.CreateUser(&user)

}

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	return s.repo.GetUserByUsername(username)
}

func hashPassword(p string) (s string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hash), nil
}
