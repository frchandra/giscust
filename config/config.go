package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	AppName          string
	IsProduction     string
	AppPort          string
	DBHost           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBPort           string
	QiscusAppId      string
	QiscusSecretKey  string
	QiscusAuthnToken string
	QiscusUrl        string
}

func getEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetAppConfig() *AppConfig {
	return initializeAppConfig()
}

func initializeAppConfig() *AppConfig {
	var appConfig = AppConfig{
		AppName:          getEnv("APP_NAME", "giscust"),
		IsProduction:     getEnv("IS_PRODUCTION", "false"),
		AppPort:          getEnv("APP_PORT", "8080"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBUser:           getEnv("DB_USER", "root"),
		DBPassword:       getEnv("DB_PASSWORD", "root"),
		DBName:           getEnv("DB_NAME", "giscust"),
		DBPort:           getEnv("DB_PORT", "5432"),
		QiscusAppId:      getEnv("QISCUS_APP_ID", ""),
		QiscusSecretKey:  getEnv("QISCUS_SECRET_KEY", ""),
		QiscusAuthnToken: getEnv("QISCUS_AUTHENTICATIOON_TOKEN", ""),
		QiscusUrl:        getEnv("QISCUS_URL", ""),
	}
	return &appConfig

}
