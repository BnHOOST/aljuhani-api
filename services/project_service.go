package services

import (
	"aljuhani-api/models"
	"aljuhani-api/repositories"
)

func GetFeaturedProjects() ([]models.Project, error) {

	return repositories.GetFeaturedProjects()
}