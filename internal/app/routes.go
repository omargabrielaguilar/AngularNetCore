package app

import (
	"my_project_go/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
}
