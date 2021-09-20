package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/xlt/edu_web/edu_ad_api/global"
	"github.com/xlt/edu_web/edu_ad_api/internal/forms"
	"github.com/xlt/edu_web/edu_ad_api/internal/proto"
	"github.com/xlt/edu_web/edu_ad_api/internal/response"
)

func SpaceSaveOrUpdate(ctx *gin.Context) {
	spaceForm := forms.SpaceForm{}
	if err := ctx.ShouldBindJSON(&spaceForm); err != nil {
		zap.S().Errorw("global.SpaceServer.GetAllSpaces failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	_, err := global.SpaceServer.SpaceSaveOrUpdate(context.Background(), &proto.SpaceSaveOrUpdateRequest{
		Id:       spaceForm.ID,
		Name:     spaceForm.Name,
		SpaceKey: spaceForm.SpaceKey,
	})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.SpaceSaveOrUpdate failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": nil,
	})
}

func GetSpaceById(ctx *gin.Context) {
	spaceID := ctx.DefaultQuery("id", "0")
	spaceIDInt, _ := strconv.Atoi(spaceID)

	rsp, err := global.SpaceServer.GetSpaceById(context.Background(), &proto.GetSpaceByIdRequest{Id: int32(spaceIDInt)})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.GetAllSpaces failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	spaceDTO := response.SpaceDTO{
		ID:         rsp.Id,
		Name:       rsp.Name,
		SpaceKey:   rsp.SpaceKey,
		CreateTime: time.Unix(rsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
		UpdateTime: time.Unix(rsp.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
		IsDel:      0,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": spaceDTO,
	})
}

func GetAllSpaces(ctx *gin.Context) {
	rsp, err := global.SpaceServer.GetAllSpaces(context.Background(), &proto.GetAllSpacesRequest{
		Pages:    0,
		PageSize: 0,
	})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.GetAllSpaces failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	SpaceDTOList := make([]response.SpaceDTO, 0)
	for _, spaceInfoRsp := range rsp.SpaceInfoResponse {
		spaceDTO := response.SpaceDTO{
			ID:         spaceInfoRsp.Id,
			Name:       spaceInfoRsp.Name,
			SpaceKey:   spaceInfoRsp.SpaceKey,
			CreateTime: time.Unix(spaceInfoRsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
			UpdateTime: time.Unix(spaceInfoRsp.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
			IsDel:      0,
		}
		SpaceDTOList = append(SpaceDTOList, spaceDTO)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": SpaceDTOList,
	})
}

func AdSaveOrUpdate(ctx *gin.Context) {
	adForm := forms.AdForm{}
	if err := ctx.ShouldBindJSON(&adForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON(&adForm) failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	loc, _ := time.LoadLocation("Local")
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", adForm.StartTime, loc)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", adForm.EndTime, loc)
	_, err := global.SpaceServer.AdSaveOrUpdate(context.Background(), &proto.AdSaveOrUpdateRequest{
		Name:        adForm.Name,
		SpaceId:     adForm.SpaceID,
		Keyword:     adForm.Keyword,
		HtmlContent: adForm.HTMLContent,
		Text:        adForm.Text,
		Link:        adForm.Link,
		StartTime:   startTime.Unix(),
		EndTime:     endTime.Unix(),
		Status:      adForm.Status,
		Priority:    adForm.Priority,
		Img:         adForm.Img,
		Id:          adForm.ID,
	})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.AdSaveOrUpdate failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": nil,
	})
}

func GetAdById(ctx *gin.Context) {
	adID := ctx.DefaultQuery("id", "0")
	adIDInt, _ := strconv.Atoi(adID)

	rsp, err := global.SpaceServer.GetAdById(context.Background(), &proto.GetAdByIdRequest{Id: int32(adIDInt)})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.GetAdById failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	adDTO := response.AdDTO{
		ID:                    int(rsp.Id),
		Name:                  rsp.Name,
		SpaceID:               int(rsp.SpaceId),
		Keyword:               rsp.Keyword,
		HTMLContent:           rsp.HtmlContent,
		Text:                  rsp.Text,
		Img:                   rsp.Img,
		Link:                  rsp.Link,
		StartTime:             time.Unix(rsp.StartTime, 0).Format("2006-01-02 15:04:05"),
		EndTime:               time.Unix(rsp.EndTime, 0).Format("2006-01-02 15:04:05"),
		CreateTime:            time.Unix(rsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
		UpdateTime:            time.Unix(rsp.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
		Status:                int(rsp.Status),
		Priority:              int(rsp.Priority),
		StartTimeFormatString: time.Unix(rsp.StartTime, 0).Format("2006-01-02 15:04:05"),
		EndTimeFormatString:   time.Unix(rsp.EndTime, 0).Format("2006-01-02 15:04:05"),
	}
	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": adDTO,
	})
}

func GetAdList(ctx *gin.Context) {
	rsp, err := global.SpaceServer.GetAdList(context.Background(), &proto.GetAdListRequest{
		Pages:    0,
		PageSize: 0,
	})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.GetAdList failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	adDTOList := make([]response.AdDTO, 0)
	for _, adInfoRsp := range rsp.AdInfoResponses {
		adDTO := response.AdDTO{
			ID:                    int(adInfoRsp.Id),
			Name:                  adInfoRsp.Name,
			SpaceID:               int(adInfoRsp.SpaceId),
			Keyword:               adInfoRsp.Keyword,
			HTMLContent:           adInfoRsp.HtmlContent,
			Text:                  adInfoRsp.Text,
			Img:                   adInfoRsp.Img,
			Link:                  adInfoRsp.Link,
			StartTime:             time.Unix(adInfoRsp.StartTime, 0).Format("2006-01-02 15:04:05"),
			EndTime:               time.Unix(adInfoRsp.EndTime, 0).Format("2006-01-02 15:04:05"),
			CreateTime:            time.Unix(adInfoRsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
			UpdateTime:            time.Unix(adInfoRsp.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
			Status:                int(adInfoRsp.Status),
			Priority:              int(adInfoRsp.Priority),
			StartTimeFormatString: time.Unix(adInfoRsp.StartTime, 0).Format("2006-01-02 15:04:05"),
			EndTimeFormatString:   time.Unix(adInfoRsp.EndTime, 0).Format("2006-01-02 15:04:05"),
		}
		adDTOList = append(adDTOList, adDTO)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": adDTOList,
	})
}

func GetAllAds(ctx *gin.Context) {
	spaceKey := ctx.Query("spaceKey")
	rsp, err := global.SpaceServer.GetAllAds(context.Background(), &proto.GetAllAdsRequest{SpaceKey: spaceKey})
	if err != nil {
		zap.S().Errorw("global.SpaceServer.GetAllAds failed", "msg", err.Error())
		HandleError(ctx, err)
		return
	}

	adDTOList := make([]response.AdDTO, 0)
	for _, adInfoResponse := range rsp.AdInfoListRsp {
		adDTO := response.AdDTO{
			ID:                    int(adInfoResponse.Id),
			Name:                  adInfoResponse.Name,
			SpaceID:               int(adInfoResponse.SpaceId),
			Keyword:               adInfoResponse.Keyword,
			HTMLContent:           adInfoResponse.HtmlContent,
			Text:                  adInfoResponse.Text,
			Img:                   adInfoResponse.Img,
			Link:                  adInfoResponse.Link,
			StartTime:             time.Unix(adInfoResponse.StartTime, 0).Format("2006-01-02 15:04:05"),
			EndTime:               time.Unix(adInfoResponse.EndTime, 0).Format("2006-01-02 15:04:05"),
			CreateTime:            time.Unix(adInfoResponse.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
			UpdateTime:            time.Unix(adInfoResponse.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
			Status:                int(adInfoResponse.Status),
			Priority:              int(adInfoResponse.Priority),
			StartTimeFormatString: time.Unix(adInfoResponse.StartTime, 0).Format("2006-01-02 15:04:05"),
			EndTimeFormatString:   time.Unix(adInfoResponse.EndTime, 0).Format("2006-01-02 15:04:05"),
		}
		adDTOList = append(adDTOList, adDTO)
	}

	spaceAdDTOSlice := make([]response.SpaceAdDTO, 0)
	spaceAdDTO := response.SpaceAdDTO{
		ID:         int(rsp.Id),
		Name:       rsp.Name,
		SpaceKey:   rsp.SpaceKey,
		CreateTime: time.Unix(rsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
		UpdateTime: time.Unix(rsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
		AdDTOList:  adDTOList,
	}
	spaceAdDTOSlice = append(spaceAdDTOSlice, spaceAdDTO)
	ctx.JSON(http.StatusOK, gin.H{
		"state":   http.StatusOK,
		"msg":     "成功",
		"content": spaceAdDTOSlice,
	})
}

//func GetAllSpaces(ctx *gin.Context) {
//	page := ctx.DefaultQuery("page", "1")
//	pageSize := ctx.DefaultQuery("pageSize", "10")
//	pageInt, _ := strconv.Atoi(page)
//	pageSizeInt, _ := strconv.Atoi(pageSize)
//
//	spaces, err := global.SpaceServer.GetAllSpaces(context.Background(), &proto.SpaceFilterRequest{
//		Pages:    int32(pageInt),
//		PageSize: int32(pageSizeInt),
//	})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.GetAllSpaces failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": spaces,
//	})
//}
//
//func GetAdBySpaceKey(ctx *gin.Context) {
//	queryArray := ctx.Query("spaceKey")
//
//	adListRsp, err := global.SpaceServer.GetAdBySpaceKey(context.Background(), &proto.SpaceKeyFilterRequest{SpaceKey: queryArray})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.GetAdBySpaceKey failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	spaceAdRsp := response.SpaceAd{
//		ID:         int(adListRsp.Id),
//		Name:       adListRsp.Name,
//		SpaceKey:   adListRsp.SpaceKey,
//		CreateTime: time.Unix(adListRsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
//		UpdateTime: time.Unix(adListRsp.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//		IsDel:      0,
//	}
//
//	adDTOList := make([]response.AdDTO, 0)
//	for _, ad := range adListRsp.SpaceAd {
//		adDTO := response.AdDTO{
//			ID:                    int(ad.Id),
//			Name:                  ad.Name,
//			SpaceID:               int(ad.SpaceId),
//			Keyword:               ad.Keyword,
//			HTMLContent:           ad.HtmlContent,
//			Text:                  ad.Text,
//			Img:                   ad.Img,
//			Link:                  ad.Link,
//			StartTime:             time.Unix(ad.StartTime, 0).Format("2006-01-02 15:04:05"),
//			EndTime:               time.Unix(ad.EndTime, 0).Format("2006-01-02 15:04:05"),
//			CreateTime:            time.Unix(ad.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
//			UpdateTime:            time.Unix(ad.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//			Status:                int(ad.Status),
//			Priority:              int(ad.Priority),
//			StartTimeFormatString: time.Unix(ad.StartTime, 0).Format("2006-01-02 15:04:05"),
//			EndTimeFormatString:   time.Unix(ad.EndTime, 0).Format("2006-01-02 15:04:05"),
//		}
//		adDTOList = append(adDTOList, adDTO)
//	}
//	spaceAdRsp.AdDTOList = adDTOList
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": spaceAdRsp,
//	})
//}
//
//func SaveOrUpdateSpace(ctx *gin.Context) {
//	spaceForm := forms.SpaceForm{}
//	if err := ctx.ShouldBindJSON(&spaceForm); err != nil {
//		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	rsp, err := global.SpaceServer.SaveOrUpdateSpace(context.Background(), &proto.SpaceInfoRequest{
//		Id:       int32(spaceForm.ID),
//		Name:     spaceForm.Name,
//		SpaceKey: spaceForm.SpaceKey,
//	})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.SaveOrUpdateSpace failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": rsp,
//	})
//}
//
//func GetSpaceById(ctx *gin.Context) {
//	spaceID := ctx.DefaultQuery("id", "0")
//	spaceIDInt, _ := strconv.Atoi(spaceID)
//
//	rsp, err := global.SpaceServer.GetSpaceById(context.Background(), &proto.SpaceByIdRequest{Id: int32(spaceIDInt)})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.SaveOrUpdateSpace failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": rsp,
//	})
//}
//
//func GetAllAds(ctx *gin.Context) {
//	page := ctx.DefaultQuery("page", "1")
//	pageSize := ctx.DefaultQuery("pageSize", "10")
//	pageInt, _ := strconv.Atoi(page)
//	pageSizeInt, _ := strconv.Atoi(pageSize)
//
//	rsp, err := global.SpaceServer.GetAllAds(context.Background(), &proto.AdPageRequest{
//		Pages:    int32(pageInt),
//		PageSize: int32(pageSizeInt),
//	})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.GetAllAds failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	adDtoList := make([]response.AdDTO, 0)
//	for _, adInfo := range rsp.SpaceAd {
//		adDTO := response.AdDTO{
//			ID:                    int(adInfo.Id),
//			Name:                  adInfo.Name,
//			SpaceID:               int(adInfo.SpaceId),
//			Keyword:               adInfo.Keyword,
//			HTMLContent:           adInfo.HtmlContent,
//			Text:                  adInfo.Text,
//			Img:                   adInfo.Img,
//			Link:                  adInfo.Link,
//			StartTime:             time.Unix(adInfo.StartTime, 0).Format("2006-01-02 15:04:05"),
//			EndTime:               time.Unix(adInfo.EndTime, 0).Format("2006-01-02 15:04:05"),
//			CreateTime:            time.Unix(adInfo.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
//			UpdateTime:            time.Unix(adInfo.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//			Status:                int(adInfo.Status),
//			Priority:              int(adInfo.Priority),
//			StartTimeFormatString: time.Unix(adInfo.StartTime, 0).Format("2006-01-02 15:04:05"),
//			EndTimeFormatString:   time.Unix(adInfo.EndTime, 0).Format("2006-01-02 15:04:05"),
//		}
//		adDtoList = append(adDtoList, adDTO)
//	}
//
//	spaceAd := response.SpaceAd{
//		ID:         int(rsp.Id),
//		Name:       rsp.Name,
//		SpaceKey:   rsp.SpaceKey,
//		CreateTime: time.Unix(rsp.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
//		UpdateTime: time.Unix(rsp.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//		AdDTOList:  adDtoList,
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": spaceAd,
//	})
//}

//func GetAllAds1(ctx *gin.Context) {
//	page := ctx.DefaultQuery("page", "1")
//	pageSize := ctx.DefaultQuery("pageSize", "10")
//	pageInt, _ := strconv.Atoi(page)
//	pageSizeInt, _ := strconv.Atoi(pageSize)
//
//	rsp, err := global.SpaceServer.GetAllAds(context.Background(), &proto.AdPageRequest{
//		Pages:    int32(pageInt),
//		PageSize: int32(pageSizeInt),
//	})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.GetAllAds failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	SpaceAd := make([]response.SpaceAd, 0)
//	for _, spaceInfo := range rsp.SpaceAdInfoListResponses {
//		adDTOList := make([]response.AdDTO, 0)
//		for _, adInfo := range spaceInfo.SpaceAd {
//			adDTO := response.AdDTO{
//				ID:                    int(adInfo.Id),
//				Name:                  adInfo.Name,
//				SpaceID:               int(adInfo.SpaceId),
//				Keyword:               adInfo.Keyword,
//				HTMLContent:           adInfo.HtmlContent,
//				Text:                  adInfo.Text,
//				Img:                   adInfo.Img,
//				Link:                  adInfo.Link,
//				StartTime:             time.Unix(adInfo.StartTime, 0).Format("2006-01-02 15:04:05"),
//				EndTime:               time.Unix(adInfo.EndTime, 0).Format("2006-01-02 15:04:05"),
//				CreateTime:            time.Unix(adInfo.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
//				UpdateTime:            time.Unix(adInfo.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//				Status:                int(adInfo.Status),
//				Priority:              int(adInfo.Priority),
//				StartTimeFormatString: time.Unix(adInfo.StartTime, 0).Format("2006-01-02 15:04:05"),
//				EndTimeFormatString:   time.Unix(adInfo.EndTime, 0).Format("2006-01-02 15:04:05"),
//			}
//			adDTOList = append(adDTOList, adDTO)
//		}
//		space := response.SpaceAd{
//			ID:         int(spaceInfo.Id),
//			Name:       spaceInfo.Name,
//			SpaceKey:   spaceInfo.SpaceKey,
//			CreateTime: time.Unix(spaceInfo.BaseProto.CreateTime, 0).Format("2006-01-02 15:04:05"),
//			UpdateTime: time.Unix(spaceInfo.BaseProto.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//			AdDTOList:  adDTOList,
//		}
//		SpaceAd = append(SpaceAd, space)
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": SpaceAd,
//	})
//}
//
//func SaveOrUpdateAd(ctx *gin.Context) {
//	adForm := forms.AdForm{}
//	if err := ctx.ShouldBindJSON(&adForm); err != nil {
//		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	loc, _ := time.LoadLocation("Local")
//	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", adForm.StartTime, loc)
//	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", adForm.EndTime, loc)
//	rsp, err := global.SpaceServer.SaveOrUpdateAd(context.Background(), &proto.AdInfoRequest{
//		Name:        adForm.Name,
//		SpaceId:     adForm.SpaceID,
//		Keyword:     adForm.Keyword,
//		HtmlContent: adForm.HtmlContent,
//		Text:        adForm.Text,
//		Link:        adForm.Link,
//		StartTime:   startTime.Unix(),
//		EndTime:     endTime.Unix(),
//		Status:      adForm.Status,
//		Priority:    adForm.Priority,
//		Img:         adForm.Img,
//		Id:          adForm.ID,
//	})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.SaveOrUpdateSpace failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state":   http.StatusInternalServerError,
//			"msg":     HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state":   http.StatusOK,
//		"msg":     "成功",
//		"content": rsp,
//	})
//}
//
//func GetAdById(ctx *gin.Context) {
//	adID := ctx.DefaultQuery("id", "0")
//	adIDInt, _ := strconv.Atoi(adID)
//
//	rsp, err := global.SpaceServer.GetAdById(context.Background(), &proto.AdByIdRequest{Id: int32(adIDInt)})
//	if err != nil {
//		zap.S().Errorw("global.SpaceServer.GetAdById failed", "msg", err.Error())
//		ctx.JSON(http.StatusOK, gin.H{
//			"state": http.StatusInternalServerError,
//			"msg": HandleError(ctx, err),
//			"content": nil,
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"state": http.StatusOK,
//		"msg": "成功",
//		"content": rsp,
//	})
//}
