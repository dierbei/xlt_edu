package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xlt/edu_web/edu_user_api/internal/router"
)

func InitRouter() *gin.Engine {
	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)
	engine := gin.Default()

	engine.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	adGroup := engine.Group("/user")
	router.InitUserRouter(adGroup)

	return engine
}
