package routes

import (
	"github.com/Edilberto-Vazquez/website-services/src/handlers"
	"github.com/Edilberto-Vazquez/website-services/src/repository"
	"github.com/gin-gonic/gin"
)

func MyInfoRoutes(repo repository.DBRepository, rg *gin.RouterGroup) {
	myinfo := rg.Group("/my-info")
	myinfo.GET("/profile", handlers.ProfileHandler(repo))
	myinfo.GET("/projects", handlers.ProjectsHandler(repo))
	myinfo.GET("/jobs", handlers.JobsHandler(repo))
}
