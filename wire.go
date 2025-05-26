//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"todolist/internal/api"
	"todolist/internal/biz"
	"todolist/internal/dao"
	"todolist/internal/service"
	"todolist/ioc"
	"todolist/pkg/logger"
)

func InitApp() *gin.Engine {
	wire.Build(
		ioc.InitDB,
		ioc.InitLogger,
		logger.NewZapLogger,
		ioc.InitGin,
		biz.NewUserBiz,
		service.NewUserService,
		dao.NewGORMUserDao,
		api.NewUserApi,
		ioc.InitRedis,
	)

	return nil
}
