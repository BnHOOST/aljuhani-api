package models

import "time"

type Project struct {
	ID uint `gorm:"primaryKey"`

	TitleAr string
	TitleEn string

	DescriptionAr string `gorm:"type:text"`
	DescriptionEn string `gorm:"type:text"`

	ImageUrl string

	IsFeatured bool
	IsActive   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}