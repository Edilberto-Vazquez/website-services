package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Edilberto-Vazquez/website-services/graph"
	"github.com/Edilberto-Vazquez/website-services/graph/generated"
	"github.com/Edilberto-Vazquez/website-services/repository"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler(repo repository.DBRepository) gin.HandlerFunc {
	resolver := graph.NewResolver(repo)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
