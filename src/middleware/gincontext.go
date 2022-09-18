package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type GinContextKey string

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxWithValue := context.WithValue(ctx.Request.Context(), "GinContextKey", ctx)
		ctx.Request = ctx.Request.WithContext(ctxWithValue)
		ctx.Next()
	}
}
