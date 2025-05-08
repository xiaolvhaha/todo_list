package ioc

import "github.com/gin-gonic/gin"

func InitGin() *gin.Engine {
	server := gin.Default()
	return server
}
