package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/edu_web/edu_ad_api/internal/api"
	"go.uber.org/zap"
)

func InitSpaceRouter(v1Group *gin.RouterGroup) {
	userGroup := v1Group.Group("/space")

	userGroup.GET("/getAllSpaces", api.GetAllSpaces)
	zap.S().Infow("初始化广告路由成功")
}
