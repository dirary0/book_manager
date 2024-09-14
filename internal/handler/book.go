package handler

import (
	"book_manager/internal/dto"
	"book_manager/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// CreateBook 创建图书
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.BookDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	if err := h.bookService.CreateBook(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "图书创建成功"})
}

// UpdateBook 更新图书
func (h *BookHandler) UpdateBook(c *gin.Context) {
	var req dto.BookDTO
	bookID := c.Param("id")
	id, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的图书ID"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	if err := h.bookService.UpdateBookByID(c.Request.Context(), uint(id), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "图书更新成功"})
}

// DeleteBook 删除图书
func (h *BookHandler) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	// Convert bookID to uint
	id, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的图书ID"})
		return
	}

	if err := h.bookService.DeleteBookByID(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "图书删除成功"})
}

// GetBookByID 通过ID获取图书
func (h *BookHandler) GetBookByID(c *gin.Context) {
	bookID := c.Param("id")

	// Convert bookID to uint
	id, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的图书ID"})
		return
	}

	book, err := h.bookService.GetBookByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// ListAllBooks 列出所有图书
func (h *BookHandler) ListAllBooks(c *gin.Context) {
	books, err := h.bookService.ListAllBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GetBookByCode 通过Code获取图书
func (h *BookHandler) GetBookByCode(c *gin.Context) {
	code := c.Param("code")

	book, err := h.bookService.GetBookByCode(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBookByCode 通过Code更新图书
func (h *BookHandler) UpdateBookByCode(c *gin.Context) {
	code := c.Param("code")
	var req dto.BookDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	if err := h.bookService.UpdateBookByCode(c.Request.Context(), code, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "图书更新成功"})
}
