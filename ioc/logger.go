package ioc

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("init zap error (%s)?", err))
	}

	zap.ReplaceGlobals(logger)

	return logger
}
