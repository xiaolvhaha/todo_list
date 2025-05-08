package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"todolist/ioc"
	"todolist/pkg/logger"
)

func InitApp() *gin.Engine {
	wire.Build(
		ioc.InitDB,
		ioc.InitLogger,
		logger.NewZapLogger,
		ioc.InitGin,
	)

	return new(gin.Engine)
}
