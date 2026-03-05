package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var App *AppConfig

type AppConfig struct {
	DatabaseDSN string
	ServerPort  string
	JWTSecret   string
}

func ConfigLoad() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Falling back to system environment variables.")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	log.Println("Environment variables loaded successfully")

	App = &AppConfig{
		DatabaseDSN: dsn,
		ServerPort:  port,
		JWTSecret:   jwtSecret,
	}

	return App
}
