package models

type Company struct {
	ID uint `gorm:"primaryKey"`

	NameAr string
	NameEn string

	SloganAr string
	SloganEn string

	AboutUsAr string `gorm:"type:text"`
	AboutUsEn string `gorm:"type:text"`

	LocationAr string
	LocationEn string

	AddressAr string
	AddressEn string

	Email string
	Phone string

	Facebook  string
	Instagram string
	LinkedIn  string
	Youtube   string
}