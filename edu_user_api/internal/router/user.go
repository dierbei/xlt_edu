package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(v1Group *gin.RouterGroup) {

	zap.S().Infow("初始化用户路由成功")
}
