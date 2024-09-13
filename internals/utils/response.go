package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func SendSuccess(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := ApiResponse{
		Success: true,
		Data:    data,
		Message: "",
	}

	json.NewEncoder(w).Encode(res)
}

func SendError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := ApiResponse{
		Success: false,
		Data:    nil,
		Message: message,
	}

	json.NewEncoder(w).Encode(res)
}
