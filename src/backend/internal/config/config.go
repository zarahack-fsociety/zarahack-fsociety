package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort  int
	DatabaseURL string
	RedisURL    string
}

func Load() (*Config, error) {
	cfg := &Config{
		ServerPort:  8080,
		DatabaseURL: "postgres://unit:unit@localhost:5432/unit?sslmode=disable",
		RedisURL:    "redis://localhost:6379/0",
	}

	if v := os.Getenv("SERVER_PORT"); v != "" {
		port, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		cfg.ServerPort = port
	}

	if v := os.Getenv("DATABASE_URL"); v != "" {
		cfg.DatabaseURL = v
	}

	if v := os.Getenv("REDIS_URL"); v != "" {
		cfg.RedisURL = v
	}

	return cfg, nil
}
