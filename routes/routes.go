package routes

import (
	"aljuhani-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api")

	{
		api.GET("/company", handlers.GetCompany)

		api.GET("/projects/featured", handlers.GetFeaturedProjects)

		api.POST(
			"/consultations",
			handlers.CreateConsultation,
		)
	}
}