package server

import (
	"book_manager/internal/handler"
	"book_manager/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(bookHandler *handler.BookHandler, userHandler handler.UserHandler) *gin.Engine {
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

	// 用户相关路由
	r.POST("/users", userHandler.CreateUser)                             // 创建用户
	r.GET("/users", userHandler.ListAllUsers)                            // 列出所有用户
	r.GET("/users/:id", userHandler.GetUserByID)                         // 通过ID获取用户
	r.PUT("/users/:id", userHandler.UpdateUser)                          // 更新用户信息
	r.DELETE("/users/:id", userHandler.DeleteUser)                       // 删除用户
	r.GET("/users/username/:username", userHandler.GetUserByUsername)    // 通过用户名获取用户
	r.PUT("/users/username/:username", userHandler.UpdateUserByUsername) // 通过用户名更新用户信息
	return r
}
