package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env            string
	Port           string
	DatabaseURL    string
	CognitoUserPool string
	CognitoClientID string
}

func LoadConfig() (*Config, error) {
	env := getEnv("ENV", "development")
	port := getEnv("PORT", "8080")
	dbURL := os.Getenv("DATABASE_URL")
	
	if dbURL == "" && env == "production" {
		return nil, fmt.Errorf("DATABASE_URL environment variable must be set in production")
	}
	
	return &Config{
		Env:             env,
		Port:            port,
		DatabaseURL:     dbURL,
		CognitoUserPool: os.Getenv("COGNITO_USER_POOL_ID"),
		CognitoClientID: os.Getenv("COGNITO_CLIENT_ID"),
	}, nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
