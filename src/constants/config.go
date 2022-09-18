package constants

import (
	"time"

	"github.com/gin-contrib/cors"
)

var (
	CorsConfig = cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://edilberto-vazquez.github.io/"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization", "Accept-Language"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	DB_NAME            = ""
	DB_COLLECTION      = ""
	ACCEPTED_LANGUAGES = map[string]string{"en-us": "en", "es-mx": "es"}
)

func SetDB(dbName, dbCollection string) {
	DB_NAME = dbName
	DB_COLLECTION = dbCollection
}
