package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Oauth    OAuthConfig
}

type AppConfig struct {
	Port int
}

type DatabaseConfig struct {
	FilePath string
}

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

var config *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Invalid PORT number: %s", err)
	}

	config = &Config{
		App: AppConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			FilePath: getEnv("DB_FILE_PATH", "test.db"),
		},
		Oauth: OAuthConfig{
			ClientID:     getEnv("OAUTH_CLIENT_ID", ""),
			ClientSecret: getEnv("OAUTH_CLIENT_SECRET", ""),
		},
	}
}

func GetConfig() *Config {
	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
