package routes

import (
	"github.com/Edilberto-Vazquez/website-services/src/constants"
	"github.com/Edilberto-Vazquez/website-services/src/handlers"
	"github.com/Edilberto-Vazquez/website-services/src/middleware"
	"github.com/Edilberto-Vazquez/website-services/src/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRoutes(s server.Server, r *gin.Engine) {
	repo := s.Repo()
	// rest api
	v1 := r.Group("/api/v1")
	v1.Use(cors.New(constants.CorsConfig))
	v1.Use(middleware.LanguageMiddleware())
	MyInfoRoutes(repo, v1)

	// graphql api
	r.Use(cors.New(constants.CorsConfig))
	r.Use(middleware.GinContextToContextMiddleware())
	r.POST("/query", handlers.GraphqlHandler(repo))
	r.GET("/playground", handlers.PlaygroundHandler())
}
