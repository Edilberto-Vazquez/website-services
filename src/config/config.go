package config

import (
	"log"
	"os"
	"strings"

	"github.com/Edilberto-Vazquez/website-services/src/constants"
	"github.com/joho/godotenv"
)

type Config struct {
	DataBaseUrl string
	Port        string
}

func getEnvVar(envVar string) (environment string) {
	switch os.Getenv("GIN_MODE") {
	case "release":
		environment = envVar
	default:
		environment = strings.Join([]string{"DEV", envVar}, "_")
	}
	return
}

func NewConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv(getEnvVar("PORT"))
	dbUrl := os.Getenv(getEnvVar("DB_URL"))
	dbName := os.Getenv(getEnvVar("DB_NAME"))
	dbCollection := os.Getenv(getEnvVar("DB_COLLECTION"))
	constants.SetDB(dbName, dbCollection)
	return Config{
		Port:        port,
		DataBaseUrl: dbUrl,
	}, nil
}
