package handlers

import (
	"encoding/json"
	"kstrategy-backend/config"
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/services"
	"kstrategy-backend/internals/utils"
	"net/http"
	"time"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var LoginInput models.LoginInput

	if err := json.NewDecoder(r.Body).Decode(&LoginInput); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := LoginInput.ValidateInput(); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.authService.Login(LoginInput)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(res.User.ID)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(config.RefreshTokenExpiry),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	utils.SendSuccess(w, res, http.StatusOK)
}
