package registry

import (
	"go.uber.org/zap"

	"github.com/xlt/edu_web/edu_ad_api/global"
)

func RegisterServer(serviceID string, freePort int) {
	if err := NewRegistry().Register(
		global.ServerConfig.Host,
		freePort,
		global.ServerConfig.Name,
		global.ServerConfig.Tags,
		serviceID,
	); err != nil {
		zap.S().Errorw("registry.NewRegistry().Register failed", "msg", err.Error())
	}
}
