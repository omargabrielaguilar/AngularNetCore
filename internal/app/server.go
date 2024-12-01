package app

import (
	"encoding/json"
	"fmt"
	"log"
	"my_project_go/internal/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	Port   string
	Router *mux.Router
}

func NewServer(configPath string) (*Server, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config struct {
		Port string `json:"port"`
	}
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %v", err)
	}

	router := mux.NewRouter()
	handlers.RegisterRoutes(router)

	return &Server{
		Port:   config.Port,
		Router: router,
	}, nil
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.Port, s.Router)
}
