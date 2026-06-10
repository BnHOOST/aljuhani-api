package handlers

import (
	"net/http"

	"aljuhani-api/dtos"
	"aljuhani-api/models"
	"aljuhani-api/services"

	"github.com/gin-gonic/gin"
)

func CreateConsultation(c *gin.Context) {

	var request dtos.CreateConsultationRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, dtos.ApiResponse{
			Success: false,
			Message: err.Error(),
		})

		return
	}

	entity := models.ConsultationRequest{
		Name: request.Name,
		Email: request.Email,
		Phone: request.Phone,
		Message: request.Message,
	}

	err := services.CreateConsultation(entity)

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			dtos.ApiResponse{
				Success: false,
				Message: err.Error(),
			},
		)

		return
	}

	c.JSON(http.StatusCreated,
		dtos.ApiResponse{
			Success: true,
			Message: "Consultation request created successfully",
		},
	)
}