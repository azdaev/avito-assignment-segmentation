package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort string
	PostgersConfig
}

type PostgersConfig struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
}

func MustLoad() Config {
	var cfg Config

	cfg.ServerPort = os.Getenv("SERVER_PORT")

	portInt, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("{config.MustLoad} error converting POSTGRES_PORT to int: %s", err)
	}
	cfg.PostgersConfig.Port = uint16(portInt)
	cfg.PostgersConfig.Host = os.Getenv("POSTGRES_HOST")
	cfg.PostgersConfig.User = os.Getenv("POSTGRES_USER")
	cfg.PostgersConfig.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.PostgersConfig.Database = os.Getenv("POSTGRES_DB")

	return cfg
}
