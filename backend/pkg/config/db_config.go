package config

import "os"

type DB_Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func DbConfigLoad() *DB_Config {
	return &DB_Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASS", "admin123"),
		DBName:     getEnv("DB_NAME", "admin123"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}
}
