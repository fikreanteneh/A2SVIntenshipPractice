package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	DatabaseURL string
	DatabaseName string
	JwtSecret   string
	JwtExpiration int
	Port string
}

func Load() (*Environment, error){
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtExpirationStr := os.Getenv("JWT_EXPIRATION")
	jwtExpiration, err := strconv.Atoi(jwtExpirationStr)
	return &Environment{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		JwtExpiration: jwtExpiration,
		Port: os.Getenv("PORT"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}, err

}