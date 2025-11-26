package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Set up environment variables
	os.Setenv("PORT", "9090")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("PACKAGES", "1,2,3")

	defer os.Unsetenv("PORT")
	defer os.Unsetenv("LOG_LEVEL")
	defer os.Unsetenv("PACKAGES")

	cfg := LoadConfig()

	expected := &Config{
		Port:            "9090",
		LogLevel:        "debug",
		PackagesDefault: []int{1, 2, 3},
	}

	if cfg.Port != expected.Port {
		t.Errorf("Expected Port: %s, got: %s", expected.Port, cfg.Port)
	}

	if cfg.LogLevel != expected.LogLevel {
		t.Errorf("Expected LogLevel: %s, got: %s", expected.LogLevel, cfg.LogLevel)
	}

	if !reflect.DeepEqual(cfg.PackagesDefault, expected.PackagesDefault) {
		t.Errorf("Expected PackagesDefault: %v, got: %v", expected.PackagesDefault, cfg.PackagesDefault)
	}
}
