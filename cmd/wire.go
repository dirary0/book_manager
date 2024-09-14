//go:build wireinject
// +build wireinject

package main

import (
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
	service.NewUserService,
	service.NewBorrowService,
)

var HandlerSet = wire.NewSet(
	handler.NewBookHandler,
	handler.NewUserHandler,
	handler.NewBorrowHandler,
)

var ServerSet = wire.NewSet(
	server.NewRouter,
)
var ConfigSet = wire.NewSet(
	repository.GetDB,
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
