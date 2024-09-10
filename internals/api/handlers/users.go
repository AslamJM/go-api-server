package handlers

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob", "Charlie"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
