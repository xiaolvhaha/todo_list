package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"todolist/internal/api"
	"todolist/pkg/middleware"
	selfValidate "todolist/pkg/middleware/validator"
)

func InitGin(userApi *api.UserApi, categoryApi *api.CategoryApi, taskApi *api.TaskApi, cache redis.Cmdable) *gin.Engine {
	server := gin.Default()
	server.Use(middleware.ValidateLogin(cache))
	RegisterValidator()
	userApi.RegisterUserRouter(server)
	categoryApi.RegisterRouter(server)
	taskApi.RegisterRouter(server)
	return server
}

func RegisterValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkEmail", selfValidate.ValidateEmail)
	}
}
