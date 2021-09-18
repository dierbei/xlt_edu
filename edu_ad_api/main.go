package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	"github.com/xlt/edu_web/edu_ad_api/global"
	"github.com/xlt/edu_web/edu_ad_api/internal/initialize"
	"github.com/xlt/edu_web/edu_ad_api/pkg/registry"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitServerClient()
	engine := initialize.InitRouter()

	if global.ServerConfig.Mode == "release" {
		// todo 线上部署需要随机端口
		//global.ServerConfig.Port = utils.GetFreePort()
	}

	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	// 注册到Consul
	registry.RegisterServer(serviceID, 8021)

	zap.S().Infow("后台服务正在启动", "port", global.ServerConfig.Port)
	go func() {
		if err := engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Fatal("后台服务启动失败")
		}
	}()

	registry.Close(serviceID, 8021)
}
