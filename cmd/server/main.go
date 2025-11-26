package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mikolabarkouski/calculator/config"
	_ "github.com/mikolabarkouski/calculator/docs"
	"github.com/mikolabarkouski/calculator/internal/api"
	"github.com/mikolabarkouski/calculator/internal/app"
	"github.com/mikolabarkouski/calculator/internal/repo"
)

func main() {
	// Try to load .env file, but don't fail if it doesn't exist (for Docker)
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: failed to load .env file: %v", err)
	}

	cfg := config.LoadConfig()

	log.Printf("Connecting to database: %s:%s/%s", cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Initialize components: repository, application, and handler
	repository, err := repo.NewRepository(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer func() {
		if err := repository.Close(); err != nil {
			log.Printf("Error closing repository: %v", err)
		}
	}()

	application := app.NewApp(repository)
	handler := api.NewHandler(application)

	log.Printf("Starting server on :%s", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handler.Router()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
