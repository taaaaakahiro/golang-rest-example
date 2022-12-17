package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port string
	Dsn  string
}

func LoadCofig() *EnvConfig {
	err := godotenv.Load(".env.sample")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cnf := &EnvConfig{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("MYSQL_DSN"),
	}

	return cnf
}
