package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

// add elem to config and matching env var
type Config struct {
	Port            string `env:"PORT" envDefault:"8080"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"info"`
	PackagesDefault []int  `env:"PACKAGES"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
