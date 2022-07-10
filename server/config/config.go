package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT              string
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
	JWT_SECRET        string
	JWT_EXPIRE_IN     string
	COOKIE_EXPIRE_IN  string
}

func LoadENV() env {
	if os.Getenv("GO_ENV") != "production" {
		// load .env file
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	loadedEnv := env{
		PORT:              os.Getenv("PORT"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
		JWT_SECRET:        os.Getenv("JWT_SECRET"),
		JWT_EXPIRE_IN:     os.Getenv("JWT_EXPIRE_IN"),
		COOKIE_EXPIRE_IN:  os.Getenv("COOKIE_EXPIRE_IN"),
	}

	return loadedEnv
}
