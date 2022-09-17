package config

import (
	"log"
	"os"

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
	log.Println(prefix)
	log.Println(port)
	log.Println(dbUrl)
	return Config{
		Port:        port,
		DataBaseUrl: dbUrl,
	}, nil
}
