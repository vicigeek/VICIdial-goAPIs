package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	// Database configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// API configuration
	APIPort string
	APIKey  string

	// Timezone
	Timezone string

	// Log level
	LogLevel string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "asterisk"),
		APIPort:    getEnv("API_PORT", "8080"),
		APIKey:     getEnv("API_KEY", ""),
		Timezone:   getEnv("TIMEZONE", "America/New_York"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv gets an environment variable with a default fallback
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
