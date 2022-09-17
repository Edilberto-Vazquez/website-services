package routes

import (
	"github.com/Edilberto-Vazquez/website-services/handlers"
	"github.com/gin-gonic/gin"
)

func MyInfoRoutes(rg *gin.RouterGroup) {
	myinfo := rg.Group("/my-info")
	myinfo.GET("/profile", handlers.ProfileHandler())
	myinfo.GET("/projects", handlers.ProjectsHandler())
	myinfo.GET("/jobs", handlers.JobsHandler())
}
