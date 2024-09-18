package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func NewValidator() {
	Validate = validator.New()
}

func (u *UserCreateInput) ValidateStruct() error {
	return Validate.Struct(u)
}

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"size:255;uniqueIndex" json:"username"`
	Email      *string   `gorm:"size:255;uniqueIndex" json:"email"`
	Password   string    `gorm:"size:255" json:"-"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserCreateInput struct {
	Username string  `json:"username" validate:"required"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password string  `json:"password" validate:"required,min=6"`
}
