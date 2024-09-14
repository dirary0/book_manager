package service

import (
	"book_manager/internal/dto"
	"book_manager/internal/models"
	"book_manager/internal/repository"
	"context"
	"errors"
	"time"
)

type BorrowService interface {
	BorrowBook(ctx context.Context, req *dto.BorrowDTO) error
	ReturnBook(ctx context.Context, borrowID uint) error
	ListUserBorrowRecords(ctx context.Context, userID uint) ([]*models.Borrow, error)
	ListUserBorrowRecordsByUsername(ctx context.Context, username string) ([]*models.Borrow, error)
}

func NewBorrowService(borrowRepo repository.BorrowRepository, bookRepo repository.BookRepository) BorrowService {
	return &borrowService{
		borrowRepository: borrowRepo,
		bookRepository:   bookRepo,
	}
}

type borrowService struct {
	borrowRepository repository.BorrowRepository
	bookRepository   repository.BookRepository
}

// BorrowBook 处理借书逻辑
func (s *borrowService) BorrowBook(ctx context.Context, req *dto.BorrowDTO) error {
	book, err := s.bookRepository.FindByID(ctx, uint(req.BookID))
	if err != nil {
		return errors.New("未找到图书: " + err.Error())
	}

	if book.Quantity <= 0 {
		return errors.New("图书库存不足")
	}

	// 创建借阅记录
	borrowRecord := &models.Borrow{
		UserID:     uint(req.UserID),
		BookID:     uint(req.BookID),
		BorrowDate: time.Now(),
	}

	if err := s.borrowRepository.Create(ctx, borrowRecord); err != nil {
		return errors.New("创建借阅记录失败: " + err.Error())
	}

	// 更新图书库存
	book.Quantity -= 1
	if err := s.bookRepository.Update(ctx, book); err != nil {
		return errors.New("更新图书库存失败: " + err.Error())
	}

	return nil
}

// ReturnBook 处理还书逻辑
func (s *borrowService) ReturnBook(ctx context.Context, borrowID uint) error {
	borrowRecord, err := s.borrowRepository.FindByID(ctx, borrowID)
	if err != nil {
		return errors.New("未找到借阅记录: " + err.Error())
	}

	if borrowRecord.ReturnDate != nil {
		return errors.New("该书已归还")
	}

	// 更新还书日期
	now := time.Now()
	borrowRecord.ReturnDate = &now
	if err := s.borrowRepository.Update(ctx, borrowRecord); err != nil {
		return errors.New("更新借阅记录失败: " + err.Error())
	}

	// 更新图书库存
	book, err := s.bookRepository.FindByID(ctx, borrowRecord.BookID)
	if err != nil {
		return errors.New("未找到图书: " + err.Error())
	}
	book.Quantity += 1
	if err := s.bookRepository.Update(ctx, book); err != nil {
		return errors.New("更新图书库存失败: " + err.Error())
	}

	return nil
}

// ListUserBorrowRecords 获取用户的借阅记录
func (s *borrowService) ListUserBorrowRecords(ctx context.Context, userID uint) ([]*models.Borrow, error) {
	records, err := s.borrowRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("获取借阅记录失败: " + err.Error())
	}
	return records, nil
}

// ListUserBorrowRecordsByUsername 获取用户的借阅记录（通过用户名）
func (s *borrowService) ListUserBorrowRecordsByUsername(ctx context.Context, username string) ([]*models.Borrow, error) {
	records, err := s.borrowRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("获取借阅记录失败: " + err.Error())
	}
	return records, nil
}
