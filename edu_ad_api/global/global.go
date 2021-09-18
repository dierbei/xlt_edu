package global

import (
	"github.com/xlt/edu_web/edu_ad_api/internal/proto"
	"gorm.io/gorm"

	"github.com/xlt/edu_web/edu_ad_api/config"
)

var (
	// Nacos
	NacosConfig = &config.NacosConfig{}

	// 全局服务配置
	ServerConfig = &config.ServerConfig{}

	// MySQL
	MySQLConn *gorm.DB

	// space server
	SpaceServer proto.SpaceClient
)
