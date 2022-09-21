package config

import (
	"log"
	"os"

	"github.com/Edilberto-Vazquez/website-services/src/constants"
	"github.com/joho/godotenv"
)

type Config struct {
	DataBaseUrl string
	Port        string
}

// func getEnvVar(envVar string) (environment string) {
// 	switch os.Getenv("GIN_MODE") {
// 	case "release":
// 		environment = envVar
// 	default:
// 		environment = strings.Join([]string{"DEV", envVar}, "_")
// 	}
// 	return
// }

func NewConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
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
