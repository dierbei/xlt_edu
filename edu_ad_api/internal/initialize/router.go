package initialize

import (
	"github.com/xlt/edu_web/edu_ad_api/internal/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)
	engine := gin.Default()

	engine.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	adGroup := engine.Group("/ad")
	router.InitSpaceRouter(adGroup)

	return engine
}
