package models

import "time"

type ConsultationRequest struct {
	ID uint `gorm:"primaryKey"`

	Name string
	Email string
	Phone string

	Message string `gorm:"type:text"`

	CreatedAt time.Time
}