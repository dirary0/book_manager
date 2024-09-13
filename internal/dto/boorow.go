package dto

type BorrowDTO struct {
	ID     uint `json:"id"`
	UserID int  `json:"user_id"`
	BookID int  `json:"book_id"`
}
