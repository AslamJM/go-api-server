package repositories

import (
	"kstrategy-backend/internals/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
