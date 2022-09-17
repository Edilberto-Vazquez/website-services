package constants

import (
	"time"

	"github.com/gin-contrib/cors"
)

const (
	DB_NAME       = "my-database"
	DB_COLLECTION = "my-info"
)

var (
	CorsConfig = cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://edilberto-vazquez.github.io/"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization", "Accept-Language"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	ACCEPTED_LANGUAGES = map[string]string{"en-us": "en", "es-mx": "es"}
)
