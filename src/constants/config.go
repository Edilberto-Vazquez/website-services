package constants

import (
	"github.com/gin-contrib/cors"
)

var (
	CorsConfig = cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://potatofy.dev", "https://goldfish-app-bixhm.ondigitalocean.app"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	}
	DB_NAME            = ""
	DB_COLLECTION      = ""
	ACCEPTED_LANGUAGES = map[string]string{"en-us": "en", "es-mx": "es"}
)

func SetDB(dbName, dbCollection string) {
	DB_NAME = dbName
	DB_COLLECTION = dbCollection
}
