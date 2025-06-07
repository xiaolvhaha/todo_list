package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"todolist/internal/api"
	"todolist/pkg/middleware"
)

func InitGin(userApi *api.UserApi, categoryApi *api.CategoryApi, taskApi *api.TaskApi, cache redis.Cmdable) *gin.Engine {
	server := gin.Default()
	server.Use(middleware.ValidateLogin(cache))
	userApi.RegisterUserRouter(server)
	categoryApi.RegisterRouter(server)
	taskApi.RegisterRouter(server)
	return server
}
