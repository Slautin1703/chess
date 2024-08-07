package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
	MongoURI      string
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		DatabaseURL:   os.Getenv("CONNECTION_STRING"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		MongoURI:      os.Getenv("MONGO_URI"),
	}

	fmt.Println(cfg)

	if cfg.DatabaseURL == "" || cfg.ServerAddress == "" {
		return nil, &MissingConfigError{}
	}

	return cfg, nil
}

type MissingConfigError struct{}

func (e *MissingConfigError) Error() string {
	return "missing required configuration"
}
