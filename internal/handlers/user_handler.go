package handlers

import (
	"encoding/json"
	"net/http"
	"my_project_go/internal/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
