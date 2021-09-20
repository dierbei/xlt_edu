package global

import (
	"github.com/xlt/edu_web/edu_ad_api/config"
	"github.com/xlt/edu_web/edu_ad_api/internal/proto"
)

var (
	// Nacos
	NacosConfig = &config.NacosConfig{}

	// 全局服务配置
	ServerConfig = &config.ServerConfig{}

	// space server
	SpaceServer proto.SpaceClient
)
