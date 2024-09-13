package main

import (
	"book_manager/config"
	"book_manager/internal/models"
	"book_manager/internal/repository"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 获取数据库连接
	db := config.GetDB()

	// 自动迁移，创建表
	err := db.AutoMigrate(&models.Book{}, &models.User{}, &models.Borrow{})
	if err != nil {
		log.Fatalf("自动创建表失败: %v", err)
	}

	log.Println("各类表创建成功.")

	// 初始化各个 repository
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db)
	borrowRepo := repository.NewBorrowRepository(db)

	// 创建测试数据
	ctx := context.Background()

	// 创建测试用户
	user := &models.User{Username: "testuser"}
	if err := userRepo.Create(ctx, user); err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
	} else {
		fmt.Printf("User created: %+v\n", user)
	}

	// 创建测试书籍
	book := &models.Book{Name: "Test Book", Quantity: 10, Code: "TESTBOOK123"}
	if err := bookRepo.Create(ctx, book); err != nil {
		fmt.Printf("Failed to create book: %v\n", err)
	} else {
		fmt.Printf("Book created: %+v\n", book)
	}

	// 创建测试借阅记录
	borrowRecord := &models.Borrow{UserID: user.ID, BookID: book.ID, BorrowDate: time.Now()}
	if err := borrowRepo.Create(ctx, borrowRecord); err != nil {
		fmt.Printf("Failed to create borrow record: %v\n", err)
	} else {
		fmt.Printf("Borrow record created: %+v\n", borrowRecord)
	}
}
