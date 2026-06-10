package dtos

type ProjectResponse struct {
	ID uint `json:"id"`

	Title string `json:"title"`

	Description string `json:"description"`

	ImageUrl string `json:"imageUrl"`

	IsFeatured bool `json:"isFeatured"`
}