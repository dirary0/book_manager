package repository

import (
	"book_manager/internal/models"
	"context"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, book *models.Book) error
	FindByID(ctx context.Context, id uint) (*models.Book, error)
	AllBooks(ctx context.Context) ([]*models.Book, error)
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

type bookRepository struct {
	db *gorm.DB
}

func (r *bookRepository) Create(ctx context.Context, book *models.Book) error {
	if err := r.db.WithContext(ctx).Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) Update(ctx context.Context, book *models.Book) error {
	if err := r.db.WithContext(ctx).Save(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) FindByID(ctx context.Context, id uint) (*models.Book, error) {
	var book models.Book
	if err := r.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) AllBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book
	if err := r.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
