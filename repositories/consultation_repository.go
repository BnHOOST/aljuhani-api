package repositories

import (
	"aljuhani-api/configs"
	"aljuhani-api/models"
)

func CreateConsultation(
	consultation models.ConsultationRequest,
) error {

	return configs.DB.
		Create(&consultation).
		Error
}