package services

import (
	"aljuhani-api/models"
	"aljuhani-api/repositories"
)

func CreateConsultation(
	consultation models.ConsultationRequest,
) error {

	return repositories.CreateConsultation(
		consultation,
	)
}