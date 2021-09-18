package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/xlt/edu_web/edu_ad_api/global"
	"github.com/xlt/edu_web/edu_ad_api/internal/proto"
)

func InitServerClient() {
	initSpaceClientLoadBalance()
}

func initSpaceClientLoadBalance() {
	spaceConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulConfig.Host,
			global.ServerConfig.ConsulConfig.Port,
			global.ServerConfig.SpaceConfig.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
	}

	spaceSrvClient := proto.NewSpaceClient(spaceConn)
	global.SpaceServer = spaceSrvClient
}
