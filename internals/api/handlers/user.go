package handlers

import (
	"encoding/json"
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/services"
	"kstrategy-backend/internals/utils"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: &service}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.GetAllUsers()

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(w, users, http.StatusOK)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.UserCreateInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := user.ValidateStruct(); err != nil {
		utils.ErrorValidateHandler(err, w)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
