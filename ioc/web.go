package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"todolist/internal/api"
	"todolist/pkg/middleware"
)

func InitGin(userApi *api.UserApi, cache redis.Cmdable) *gin.Engine {
	server := gin.Default()
	server.Use(middleware.ValidateLogin(cache))
	userApi.RegisterUserRouter(server)
	return server
}
