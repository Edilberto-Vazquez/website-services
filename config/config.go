package config

import (
	"log"
	"os"

	"github.com/Edilberto-Vazquez/website-services/constants"
	"github.com/joho/godotenv"
)

type Config struct {
	DataBaseUrl string
	Port        string
}

func SetEnvironment() (environment string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	switch os.Getenv("APP_ENV") {
	case "production":
		environment = ""
	default:
		environment = "DEV_"
	}
	return
}

func NewConfig() (Config, error) {
	prefix := SetEnvironment()
	port := os.Getenv(prefix + "PORT")
	dbUrl := os.Getenv(prefix + "DB_URL")
	dbName := os.Getenv(prefix + "DB_NAME")
	dbCollection := os.Getenv(prefix + "DB_COLLECTION")
	constants.SetDB(dbName, dbCollection)
	return Config{
		Port:        port,
		DataBaseUrl: dbUrl,
	}, nil
}
