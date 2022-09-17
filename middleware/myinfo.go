package middleware

import (
	"net/http"

	"github.com/Edilberto-Vazquez/website-services/i18n"
	"github.com/gin-gonic/gin"
)

func LanguageMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.GetHeader("Accept-Language")
		value, exists := i18n.CheckLanguage(lang)
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
