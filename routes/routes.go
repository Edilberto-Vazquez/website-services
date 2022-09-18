package routes

import (
	"github.com/Edilberto-Vazquez/website-services/constants"
	"github.com/Edilberto-Vazquez/website-services/middleware"
	"github.com/Edilberto-Vazquez/website-services/server"
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
	gql := r.Group("/gql")
	gql.Use(middleware.LanguageMiddleware())
	gql.Use(middleware.GinContextToContextMiddleware())
	GqlRoutes(repo, gql)

}
