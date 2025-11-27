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
