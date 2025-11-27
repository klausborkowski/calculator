package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/klausborkowski/calculator/config"
	_ "github.com/klausborkowski/calculator/docs"
	"github.com/klausborkowski/calculator/internal/api"
	"github.com/klausborkowski/calculator/internal/app"
	"github.com/klausborkowski/calculator/internal/repo"
)

func main() {
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Panic("env load err")
	}

	cfg := config.LoadConfig()

	log.Printf("PACKAGES DEFAULT:%v", cfg.PackagesDefault)

	//define components -> repo,app and handler(api)
	repository := repo.NewRepository(cfg.PackagesDefault)
	application := app.NewApp(repository)
	handler := api.NewHandler(application)

	log.Printf("Starting server on :%s\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handler.Router()); err != nil {
		log.Fatal("Server failed:", err)
	}
}
