package config

import (
	"os"

	"github.com/Edilberto-Vazquez/website-services/src/constants"
)

type Config struct {
	DataBaseUrl string
	Port        string
}

func NewConfig() (Config, error) {
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbCollection := os.Getenv("DB_COLLECTION")
	constants.SetDB(dbName, dbCollection)
	return Config{
		Port:        port,
		DataBaseUrl: dbUrl,
	}, nil
}
