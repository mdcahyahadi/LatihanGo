package uploadphoto

import (
	"latihan/infrastructure/appLog"
	"latihan/infrastructure/configuration"
	"latihan/shared/response"
	"latihan/shared/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

type Service struct {
	conf   *configuration.AppConfig
	appLog *appLog.AppLog
	repo   IRepository
}

func NewService(conf *configuration.AppConfig, appLog *appLog.AppLog, repo IRepository) *Service {
	return &Service{conf: conf, appLog: appLog, repo: repo}
}

func (svc *Service) GetListPhoto(ctx *gin.Context, req GetListPhoto) (resp []PhotoEntity, err error) {
	userID := ctx.GetHeader(utils.MyGramHeaderUserID)
	span, apmCtx := apm.StartSpan(ctx.Request.Context(), "GetListPhoto", response.APMSpanTypeLogicalProcess)
	defer span.End()
	resp, err = svc.repo.GetPhotos(apmCtx, req.Limit, req.SearchBy, req.Keyword)
	if err != nil {
		svc.appLog.LogWarn(userID, "error getting photos", err)
		return
	}
	return
}

func (svc *Service) PostingPhoto(ctx *gin.Context, req SavePhoto) (resp interface{}, err error) {
	userID, _ := strconv.Atoi(ctx.GetHeader(utils.MyGramHeaderUserID))
	createTime := time.Now()
	span, apmCtx := apm.StartSpan(ctx.Request.Context(), "PostingPhoto", response.APMSpanTypeLogicalProcess)
	defer span.End()
	resp, err = svc.repo.StorePhoto(apmCtx, PhotoEntity{
		Title:     req.Title,
		Caption:   req.Caption,
		URL:       req.URL,
		UserID:    int64(userID),
		CreatedAt: createTime,
		UpdatedAt: createTime,
	})
	if err != nil {
		svc.appLog.LogWarn(strconv.Itoa(userID), "error posting photo", err)
		return
	}
	return
}

func (svc *Service) EditPhoto(ctx *gin.Context, req EditPhoto) (resp interface{}, err error) {
	userID, _ := strconv.Atoi(ctx.GetHeader(utils.MyGramHeaderUserID))
	updateTime := time.Now()
	span, apmCtx := apm.StartSpan(ctx.Request.Context(), "EditPhoto", response.APMSpanTypeLogicalProcess)
	defer span.End()
	err = svc.repo.UpdatePhoto(apmCtx, PhotoEntity{
		ID:        req.PhotoID,
		Title:     req.NewTitle,
		Caption:   req.NewCaption,
		UpdatedAt: updateTime,
	})
	if err != nil {
		svc.appLog.LogWarn(strconv.Itoa(userID), "error editing photo", err)
		return
	}
	return
}
