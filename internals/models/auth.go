package models

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string
	User        User
}

func (l *LoginInput) ValidateInput() error {
	return Validate.Struct(l)
}
