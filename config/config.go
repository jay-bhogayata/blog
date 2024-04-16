package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/jay-bhogayata/blogapi/logger"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DBURL string
	}
	Env         string
	EmailSender string
	JWTSecret   string
}

func LoadConfig() (*Config, error) {

	var cfg Config

	cfg.Server.Port = os.Getenv("SERVER_PORT")
	if cfg.Server.Port == "" {
		logger.Log.Warn("no SERVER_PORT env variable provided defaulting to port 8080")
		cfg.Server.Port = "8080"
	}
	if _, err := strconv.Atoi(cfg.Server.Port); err != nil {
		logger.Log.Warn("invalid SERVER_PORT, using default port 8080")

		cfg.Server.Port = "8080"
	}

	cfg.Database.DBURL = os.Getenv("DATABASE_URL")
	if cfg.Database.DBURL == "" {
		logger.Log.Error("no DATABASE_URL found in env file")
		return nil, errors.New("DATABASE_URL env not found")
	}

	cfg.Env = os.Getenv("ENV")
	if cfg.Env == "" {
		logger.Log.Error("no ENV env variable provided defaulting to dev")
		cfg.Env = "DEV"
	}

	cfg.EmailSender = os.Getenv("MAILER_SENDER")
	if cfg.EmailSender == "" {
		logger.Log.Error("no MAILER_SENDER env variable provided")
		return nil, errors.New("MAILER_SENDER env not found")
	}

	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	if cfg.JWTSecret == "" {
		logger.Log.Error("no JWT_SECRET env variable provided")
		return nil, errors.New("JWT_SECRET env not found")
	}

	return &cfg, nil
}
