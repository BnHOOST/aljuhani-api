package dtos

type CompanyResponse struct {
	Name     string `json:"name"`
	Slogan   string `json:"slogan"`
	AboutUs  string `json:"aboutUs"`
	Location string `json:"location"`
	Address  string `json:"address"`

	Email string `json:"email"`
	Phone string `json:"phone"`

	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	LinkedIn  string `json:"linkedin"`
	Youtube   string `json:"youtube"`
}