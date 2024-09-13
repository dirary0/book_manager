package server

import (
	"book_manager/internal/handler"
	"book_manager/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(bookHandler *handler.BookHandler) *gin.Engine {
	r := gin.Default()
	r.Use(
		middleware.CORS(),
	)
	// 图书相关路由
	r.POST("/books", bookHandler.CreateBook)                 // 创建图书
	r.GET("/books", bookHandler.ListAllBooks)                // 列出所有图书
	r.GET("/books/:id", bookHandler.GetBookByID)             // 通过ID获取图书
	r.PUT("/books/:id", bookHandler.UpdateBook)              // 更新图书
	r.DELETE("/books/:id", bookHandler.DeleteBook)           // 删除图书
	r.GET("/books/code/:code", bookHandler.GetBookByCode)    // 通过Code获取图书
	r.PUT("/books/code/:code", bookHandler.UpdateBookByCode) // 通过Code更新图书

	return r
}
