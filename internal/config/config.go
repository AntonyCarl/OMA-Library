package config

import (
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func SetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal(err)
	}
	var cfg Config

	err = envconfig.Process("db", &cfg.DB)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	return &cfg
}
