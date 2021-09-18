package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/xlt/edu_web/edu_ad_api/global"
	"github.com/xlt/edu_web/edu_ad_api/internal/proto"
)

func GetAllSpaces(ctx *gin.Context) {
	spaces, err := global.SpaceServer.GetAllSpaces(context.Background(), &proto.SpaceFilterRequest{
		Pages:    0,
		PageSize: 0,
	})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.GetAllSpaces failed", "msg", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": spaces,
	})
}
