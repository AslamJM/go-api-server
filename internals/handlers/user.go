package handlers

import (
	"encoding/json"
	"kstrategy-backend/internals/services"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: &service}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.service.GetAllUsers()

	if err != nil {
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
