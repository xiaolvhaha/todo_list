package main

import (
	biz "todolist/internal/biz"
	"todolist/internal/dao"
	"todolist/internal/service"
	"todolist/ioc"
	"todolist/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() *gin.Engine {
	wire.Build(
		ioc.InitDB,
		ioc.InitLogger,
		logger.NewZapLogger,
		ioc.InitGin,
		dao.NewGORMUserDao,
		biz.NewUserBiz,
		service.NewUserService,
		wire.Bind(new(biz.UserDao), new(*dao.GORMUserDao)),
		wire.Bind(new(service.UserBiz), new(*biz.UserUsecase)),
	)

	return ioc.InitGin()
}
