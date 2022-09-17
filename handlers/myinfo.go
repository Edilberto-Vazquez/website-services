package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		ctx.JSON(http.StatusOK, lang)
	}
}

func ProjectsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		ctx.JSON(http.StatusOK, lang)
	}
}

func ResumeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		ctx.JSON(http.StatusOK, lang)
	}
}

func JobsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		ctx.JSON(http.StatusOK, lang)
	}
}
