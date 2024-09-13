package main

import (
	"book_manager/config"
	"book_manager/internal/models"
	"log"
)

func main() {
	config.InitDB()

	db := config.GetDB()

	// 自动迁移，创建表
	err := db.AutoMigrate(&models.Book{}, &models.User{}, &models.Borrow{})
	if err != nil {
		log.Fatalf("自动创建表失败: %v", err)
	}

	log.Println("各类表创建成功.")

	//启动 HTTP 服务或其他应用逻辑
	//r := server.NewRouter(...)
	//r.Run()
}
