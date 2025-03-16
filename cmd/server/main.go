package main

import (
	"log"
	"task-manager/internal/database"
	"task-manager/internal/routes"
)

func main() {
	if err := database.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.SetupRouter()

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}