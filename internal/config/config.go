package config

import (
	"log"
	"os"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	PostgersConfig
}

type PostgersConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     uint16 `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"admin"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	Database string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("{config.MustLoad} CONFIG_PATH env-var is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("{config.MustLoad} error opening config file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("{config.MustLoad} error reading config file: %s", err)
	}

	cfg.Se

	portInt, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("{config.MustLoad} error converting POSTGRES_PORT to int: %s", err)
	}
	cfg.PostgersConfig.Port = uint16(portInt)
	cfg.PostgersConfig.Host = os.Getenv("POSTGRES_HOST")
	cfg.PostgersConfig.User = os.Getenv("POSTGRES_USER")
	cfg.PostgersConfig.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.PostgersConfig.Database = os.Getenv("POSTGRES_DATABASE")

	return cfg
}
