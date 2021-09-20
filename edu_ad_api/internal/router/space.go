package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/edu_web/edu_ad_api/internal/api"
	"go.uber.org/zap"
)

func InitSpaceRouter(v1Group *gin.RouterGroup) {
	//v1Group.GET("/getAdBySpaceKey", api.GetAdBySpaceKey)
	v1Group.GET("/getAllAds", api.GetAllAds)
	v1Group.GET("/getAdList", api.GetAdList)
	v1Group.POST("/saveOrUpdate", api.AdSaveOrUpdate)
	v1Group.GET("/getAdById", api.GetAdById)

	spaceGroup := v1Group.Group("/space")
	spaceGroup.GET("/getAllSpaces", api.GetAllSpaces)
	spaceGroup.POST("/saveOrUpdate", api.SpaceSaveOrUpdate)
	spaceGroup.GET("/getSpaceById", api.GetSpaceById)

	zap.S().Infow("初始化广告路由成功")
}
