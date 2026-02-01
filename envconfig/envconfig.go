package envconfig

import (
	"log"
	"os"
)

type EnvConfig struct {
	DATABASE_URL   string
	TestDBname     string
	TestDBhost     string
	GoogleClientID string
}

var AppConfig EnvConfig

func LoadEnvConfig() {
	AppConfig = EnvConfig{
		DATABASE_URL:   os.Getenv("DATABASE_URL"),
		TestDBname:     os.Getenv("TEST_DB_NAME"),
		TestDBhost:     os.Getenv("TEST_DB_HOST"),
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
	}

	if AppConfig.GoogleClientID == "" {
		log.Fatal("Error: GoogleClientID must be set")
	}
}
