package middleware

import (
	"net/http"

	"github.com/Edilberto-Vazquez/website-services/src/i18n"
	"github.com/gin-gonic/gin"
)

type Language struct {
	Lang string `form:"lang" binding:"required"`
}

func LanguageMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var lang Language
		if ctx.ShouldBindQuery(&lang) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "language must be provided",
			})
			ctx.Abort()
			return
		}
		value, exists := i18n.CheckLanguage(lang.Lang)
		if !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "The app cannot provide that language",
			})
			ctx.Abort()
			return
		}
		ctx.Set("lang", value)
		ctx.Next()
	}
}
