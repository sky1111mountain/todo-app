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
		DBName:         os.Getenv("DATABASE"),
		DBHost:         os.Getenv("DBHOST"),
		TestDBname:     os.Getenv("TESTDB"),
		TestDBhost:     os.Getenv("TESTDBHOST"),
		GoogleClientID: os.Getenv("GOOGLECLIENTID"),
	}

	if AppConfig.GoogleClientID == "" {
		log.Fatal("Error: GoogleClientID must be set")
	}
}
