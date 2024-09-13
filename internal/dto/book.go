package dto

type BookDTO struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Code     string `json:"code"`
}
