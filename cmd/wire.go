//go:build wireinject
// +build wireinject

package main

import (
	"book_manager/config"
	"book_manager/internal/handler"
	"book_manager/internal/repository"
	"book_manager/internal/server"
	"book_manager/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewBookRepository,
	repository.NewUserRepository,
	repository.NewBorrowRepository,
)

var ServiceSet = wire.NewSet(
	service.NewBookService,
)

var HandlerSet = wire.NewSet(
	handler.NewBookHandler,
)

var ServerSet = wire.NewSet(
	server.NewRouter,
)
var ConfigSet = wire.NewSet(
	config.GetDB,
)

func newApp() (*gin.Engine, func(), error) {
	wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		ConfigSet,
	)
	return &gin.Engine{}, nil, nil
}
