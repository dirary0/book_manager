package main

import (
	"book_manager/internal/repository"
	"log"
	"os"
)

func main() {
	repository.InitDB()
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	app, cleanup, err := newApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动 HTTP 服务
	if err := app.Run(); err != nil {
		logger.Fatalf("启动失败: %v", err)
	}
}
