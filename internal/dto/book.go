package dto

type BookDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Code     string `json:"code"`
}
