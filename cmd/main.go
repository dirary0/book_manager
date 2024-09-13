package main

import (
	"book_manager/config"
	"log"
	"os"
)

func main() {
	config.InitDB()

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	// 通过 Wire 自动生成的依赖注入代码创建应用程序
	app, cleanup, err := newApp(logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动 HTTP 服务
	if err := app.Run(); err != nil {
		logger.Fatalf("启动失败: %v", err)
	}
}
