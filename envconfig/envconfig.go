package envconfig

import (
	"log"
	"os"
)

type EnvConfig struct {
	DBUser         string
	DBPass         string
	DBName         string
	DBHost         string
	TestDBname     string
	TestDBhost     string
	GoogleClientID string
}

var AppConfig EnvConfig

func LoadEnvConfig() {
	AppConfig = EnvConfig{
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASS"),
		DBName:         os.Getenv("DB_NAME"),
		DBHost:         os.Getenv("DB_HOST"),
		TestDBname:     os.Getenv("TEST_DB_NAME"),
		TestDBhost:     os.Getenv("TEST_DB_HOST"),
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
	}

	if AppConfig.GoogleClientID == "" {
		log.Fatal("Error: GoogleClientID must be set")
	}
}
