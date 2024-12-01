package main

import (
	"fmt"
	"log"
	"my_project_go/internal/app"
	"os"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/config.json"
	}

	server, err := app.NewServer(configPath)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	fmt.Printf("Starting server on port %s...\n", server.Port)
	log.Fatal(server.Start())
}
