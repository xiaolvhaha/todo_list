package ioc

import (
	"github.com/gin-gonic/gin"
	"todolist/internal/api"
	"todolist/pkg/middleware"
)

func InitGin(userApi *api.UserApi) *gin.Engine {
	server := gin.Default()
	server.Use(middleware.ValidateLogin())
	userApi.RegisterUserRouter(server)
	return server
}
