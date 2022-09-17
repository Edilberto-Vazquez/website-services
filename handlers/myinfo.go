package handlers

import (
	"net/http"

	"github.com/Edilberto-Vazquez/website-services/repository"
	"github.com/gin-gonic/gin"
)

func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		profile, err := repository.GetProfile(ctx, lang)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, profile)
	}
}

func ProjectsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		profile, err := repository.GetProjects(ctx, lang)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, profile)
	}
}

func JobsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		profile, err := repository.GetJobs(ctx, lang)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, profile)
	}
}
