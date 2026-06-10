package repositories

import (
	"aljuhani-api/configs"
	"aljuhani-api/models"
)

func GetFeaturedProjects() ([]models.Project, error) {

	var projects []models.Project

	err := configs.DB.
		Where("is_featured = ? AND is_active = ?", true, true).
		Order("created_at desc").
		Find(&projects).Error

	return projects, err
}