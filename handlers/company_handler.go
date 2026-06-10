package handlers

import (
	"net/http"

	"aljuhani-api/configs"
	"aljuhani-api/dtos"
	"aljuhani-api/models"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {

	var company models.Company

	result := configs.DB.First(&company)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Company information not found",
		})

		return
	}

	lang := c.GetString("lang")

	response := dtos.CompanyResponse{
		Email:     company.Email,
		Phone:     company.Phone,
		Facebook:  company.Facebook,
		Instagram: company.Instagram,
		LinkedIn:  company.LinkedIn,
		Youtube:   company.Youtube,
	}

	if lang == "ar" {

		response.Name = company.NameAr
		response.Slogan = company.SloganAr
		response.AboutUs = company.AboutUsAr
		response.Location = company.LocationAr
		response.Address = company.AddressAr

	} else {

		response.Name = company.NameEn
		response.Slogan = company.SloganEn
		response.AboutUs = company.AboutUsEn
		response.Location = company.LocationEn
		response.Address = company.AddressEn
	}

	c.JSON(http.StatusOK, dtos.ApiResponse{
		Success: true,
		Message: "Success",
		Data:    response,
	})
}