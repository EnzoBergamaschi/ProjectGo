package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env                string
	AppPort            string
	DBUser             string
	DBPass             string
	DBHost             string
	DBPort             string
	DBName             string
	JWTSecret          string
	JWTExpirationHours int
}

func Load() *Config {
	// Carrega o .env se existir
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis do ambiente")
	}

	jwtExp := 24
	if v := os.Getenv("JWT_EXP_HOURS"); v != "" {
		// opcional: parse int com strconv.Atoi
	}

	return &Config{
		Env:                getEnv("ENV", "development"),
		AppPort:            getEnv("APP_PORT", "8080"),
		DBUser:             getEnv("DB_USER", "root"),
		DBPass:             getEnv("DB_PASS", ""),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "3306"),
		DBName:             getEnv("DB_NAME", "projectgo"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTExpirationHours: jwtExp,
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
