package routes

import (
	"github.com/Edilberto-Vazquez/website-services/handlers"
	"github.com/Edilberto-Vazquez/website-services/repository"
	"github.com/gin-gonic/gin"
)

func GqlRoutes(repo repository.DBRepository, rg *gin.RouterGroup) {
	rg.POST("/query", handlers.GraphqlHandler(repo))
	rg.GET("/", handlers.PlaygroundHandler())
}
