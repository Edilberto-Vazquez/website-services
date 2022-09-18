package routes

import (
	"github.com/Edilberto-Vazquez/website-services/handlers"
	"github.com/gin-gonic/gin"
)

func GqlRoutes(rg *gin.RouterGroup) {
	rg.POST("/query", handlers.GraphqlHandler())
	rg.GET("/", handlers.PlaygroundHandler())
}
