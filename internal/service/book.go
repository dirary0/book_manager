package service

import (
	"book_manager/internal/dto"
	"book_manager/internal/models"
	"book_manager/internal/repository"
	"context"
	"errors"
)

type BookService interface {
	CreateBook(ctx context.Context, req *dto.BookDTO) error
	UpdateBookByID(ctx context.Context, bookID uint, req *dto.BookDTO) error
	DeleteBookByID(ctx context.Context, bookID uint) error
	GetBookByID(ctx context.Context, bookID uint) (*models.Book, error)
	ListAllBooks(ctx context.Context) ([]*models.Book, error)
	GetBookByCode(ctx context.Context, code string) (*models.Book, error)
	UpdateBookByCode(ctx context.Context, code string, req *dto.BookDTO) error
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

type bookService struct {
	bookRepository repository.BookRepository
}

// CreateBook 创建图书
func (s *bookService) CreateBook(ctx context.Context, req *dto.BookDTO) error {
	book := &models.Book{
		Name:     req.Name,
		Quantity: req.Quantity,
		Code:     req.Code,
	}
	if err := s.bookRepository.Create(ctx, book); err != nil {
		return errors.New("创建图书失败: " + err.Error())
	}
	return nil
}

// UpdateBookByID 更新图书信息
func (s *bookService) UpdateBookByID(ctx context.Context, bookID uint, req *dto.BookDTO) error {
	book, err := s.bookRepository.FindByID(ctx, bookID)
	if err != nil {
		return errors.New("未找到图书: " + err.Error())
	}
	book.Name = req.Name
	book.Quantity = req.Quantity
	book.Code = req.Code
	if err := s.bookRepository.Update(ctx, book); err != nil {
		return errors.New("更新图书失败: " + err.Error())
	}
	return nil
}

// DeleteBookByID 删除图书
func (s *bookService) DeleteBookByID(ctx context.Context, bookID uint) error {
	if err := s.bookRepository.Delete(ctx, bookID); err != nil {
		return errors.New("删除图书失败: " + err.Error())
	}
	return nil
}

// GetBookByID 通过ID获取图书
func (s *bookService) GetBookByID(ctx context.Context, bookID uint) (*models.Book, error) {
	book, err := s.bookRepository.FindByID(ctx, bookID)
	if err != nil {
		return nil, errors.New("未找到图书: " + err.Error())
	}
	return book, nil
}

// ListAllBooks 列出所有图书
func (s *bookService) ListAllBooks(ctx context.Context) ([]*models.Book, error) {
	books, err := s.bookRepository.AllBooks(ctx)
	if err != nil {
		return nil, errors.New("获取图书列表失败: " + err.Error())
	}
	return books, nil
}

// GetBookByCode 通过Code查询书籍
func (s *bookService) GetBookByCode(ctx context.Context, code string) (*models.Book, error) {
	book, err := s.bookRepository.FindByCode(ctx, code)
	if err != nil {
		return nil, errors.New("未找到图书: " + err.Error())
	}
	return book, nil
}

// UpdateBookByCode 通过Code更新书籍
func (s *bookService) UpdateBookByCode(ctx context.Context, code string, req *dto.BookDTO) error {
	book, err := s.bookRepository.FindByCode(ctx, code)
	if err != nil {
		return errors.New("未找到图书: " + err.Error())
	}

	// 更新书籍信息
	book.Name = req.Name
	book.Quantity = req.Quantity
	book.Code = req.Code

	if err := s.bookRepository.UpdateByCode(ctx, code, book); err != nil {
		return errors.New("更新图书失败: " + err.Error())
	}
	return nil
}
