package handlers

import (
	"fmt"
	"net/http"

	"aljuhani-api/dtos"
	"aljuhani-api/helpers"
	"aljuhani-api/services"

	"github.com/gin-gonic/gin"
)

func GetFeaturedProjects(c *gin.Context) {

	lang := c.GetString("lang")

	projects, err := services.GetFeaturedProjects()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if len(projects) == 0 {
		c.JSON(http.StatusOK, dtos.ApiResponse{
			Success: true,
			Message: "No Projects Found Yet!",
			Data:    nil,
		})
		return
	}


	var response []dtos.ProjectResponse

	for _, p := range projects {

		project := dtos.ProjectResponse{
			ID:         p.ID,
			ImageUrl:   BuildImageURL(c, p.ImageUrl),
			IsFeatured: p.IsFeatured,
		}

		if helpers.IsArabic(lang) {
			project.Title = p.TitleAr
			project.Description = p.DescriptionAr
		} else {
			project.Title = p.TitleEn
			project.Description = p.DescriptionEn
		}

		response = append(response, project)
	}

	c.JSON(http.StatusOK, dtos.ApiResponse{
		Success: true,
		Message: "Success",
		Data:    response,
	})
}

// Suppose you have a helper function:
func BuildImageURL(c *gin.Context, filename string) string {
    if filename == "" {
        return ""
    }
    // Get scheme (http/https) and host from the request
    scheme := "http"
    if c.Request.TLS != nil {
        scheme = "https"
    }
    host := c.Request.Host

    // Construct full URL
    return fmt.Sprintf("%s://%s/public/imgs/%s", scheme, host, filename)
}
