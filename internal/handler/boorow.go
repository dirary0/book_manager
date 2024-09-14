package handler

import (
	"book_manager/internal/dto"
	"book_manager/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BorrowHandler struct {
	borrowService service.BorrowService
}

func NewBorrowHandler(borrowService service.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		borrowService: borrowService,
	}
}

// BorrowBook 处理借书请求
func (h *BorrowHandler) BorrowBook(c *gin.Context) {
	var req dto.BorrowDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	if err := h.borrowService.BorrowBook(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "借书成功"})
}

// ReturnBook 处理还书请求
func (h *BorrowHandler) ReturnBook(c *gin.Context) {
	borrowIDStr := c.Param("id")
	borrowID, err := strconv.ParseUint(borrowIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的借阅记录ID"})
		return
	}

	if err := h.borrowService.ReturnBook(c.Request.Context(), uint(borrowID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "还书成功"})
}

// ListUserBorrowRecordsById 获取用户的借阅记录
func (h *BorrowHandler) ListUserBorrowRecordsById(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	records, err := h.borrowService.ListUserBorrowRecords(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}

// ListUserBorrowRecordsByName 获取用户的借阅记录（通过用户名）
func (h *BorrowHandler) ListUserBorrowRecordsByName(c *gin.Context) {
	username := c.Param("user_name")

	records, err := h.borrowService.ListUserBorrowRecordsByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}
