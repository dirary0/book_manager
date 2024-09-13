package repository

import (
	"book_manager/internal/models"
	"context"
	"gorm.io/gorm"
)

type BorrowRepository interface {
	Create(ctx context.Context, record *models.Borrow) error
	Update(ctx context.Context, record *models.Borrow) error
	FindByID(ctx context.Context, id uint) (*models.Borrow, error)
	AllBorrowRecords(ctx context.Context) ([]*models.Borrow, error)
	FindByUserID(ctx context.Context, userID uint) ([]*models.Borrow, error)
	FindByUsername(ctx context.Context, username string) ([]*models.Borrow, error)
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowRepository{
		db: db,
	}
}

type borrowRepository struct {
	db *gorm.DB
}

func (r *borrowRepository) Create(ctx context.Context, record *models.Borrow) error {
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	return nil
}

func (r *borrowRepository) Update(ctx context.Context, record *models.Borrow) error {
	if err := r.db.WithContext(ctx).Save(record).Error; err != nil {
		return err
	}
	return nil
}

func (r *borrowRepository) FindByID(ctx context.Context, id uint) (*models.Borrow, error) {
	var record models.Borrow
	if err := r.db.WithContext(ctx).First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *borrowRepository) AllBorrowRecords(ctx context.Context) ([]*models.Borrow, error) {
	var records []*models.Borrow
	if err := r.db.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *borrowRepository) FindByUserID(ctx context.Context, userID uint) ([]*models.Borrow, error) {
	var records []*models.Borrow
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
func (r *borrowRepository) FindByUsername(ctx context.Context, username string) ([]*models.Borrow, error) {
	var records []*models.Borrow
	if err := r.db.WithContext(ctx).Joins("JOIN users ON users.id = borrows.user_id").Where("users.username = ?", username).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
