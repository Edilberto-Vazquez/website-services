package handlers

import (
	"net/http"

	"github.com/Edilberto-Vazquez/website-services/src/repository"
	"github.com/gin-gonic/gin"
)

func ProfileHandler(repo repository.DBRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		profile, err := repo.GetProfile(ctx, lang)
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

func ProjectsHandler(repo repository.DBRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		profile, err := repo.GetProjects(ctx, lang)
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

func JobsHandler(repo repository.DBRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		profile, err := repo.GetJobs(ctx, lang)
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

func FullProfileHandler(repo repository.DBRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.MustGet("lang").(string)
		fullProfile, err := repo.GetFullProfile(ctx, lang)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, fullProfile)
	}
}
