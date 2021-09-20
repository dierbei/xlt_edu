package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
)

func HandleError(ctx *gin.Context, err error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			ctx.JSON(http.StatusOK, gin.H{
				"state":   http.StatusInternalServerError,
				"msg":     e.Message(),
				"content": nil,
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusInternalServerError,
		"msg":     err.Error(),
		"content": nil,
	})
}
