package uploadphoto

import (
	"latihan/infrastructure/activityLog"
	"latihan/infrastructure/appLog"
	"latihan/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

type Delivery struct {
	appLog      *appLog.AppLog
	service     IService
	activityLog activityLog.IService
}

func NewDelivery(appLog *appLog.AppLog, service IService, activityLog activityLog.IService) *Delivery {
	return &Delivery{appLog: appLog, service: service, activityLog: activityLog}
}

func (d *Delivery) GetPhotos(ctx *gin.Context) {
	span, apmCtx := apm.StartSpan(ctx.Request.Context(), "GetPhotos", response.APMSpanTypeHandlerProcess)
	defer span.End()
	trx := apm.TransactionFromContext(apmCtx)

	var req GetListPhoto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		d.appLog.LogResponse(http.StatusBadRequest, ctx.Request, req, response.ErrBadRequest)
		trx.Context.SetCustom(response.APMTrxCustomKeyResponse, response.ErrBadRequest)
		response.SendErrorResponse4xx(ctx, http.StatusBadRequest, response.ErrBadRequest)
		return
	}

	d.appLog.LogRequest(ctx.Request, req)
	trx.Context.SetCustom(response.APMTrxCustomKeyRequestParam, req)

	data, err := d.service.GetListPhoto(ctx, req)
	if err != nil {
		switch err.Error() {
		case ErrUnAuthorized.Error():
			d.appLog.LogResponse(http.StatusUnauthorized, ctx.Request, req, ErrUnAuthorized)
			trx.Context.SetCustom(response.APMTrxCustomKeyResponse, ErrUnAuthorized)
			response.SendErrorResponse4xx(ctx, http.StatusUnauthorized, ErrUnAuthorized)
			return
		default:
			resp := response.ErrorStruct{
				Code:        response.RCErrorGeneral,
				Description: response.DescriptionFailed,
				Message:     err.Error(),
			}
			d.appLog.LogResponse(http.StatusInternalServerError, ctx.Request, req, resp)
			trx.Context.SetCustom(response.APMTrxCustomKeyResponse, resp)
			response.SendErrorResponse5xx(ctx, resp)
			return
		}
	}

	resp := response.Response{
		ResponseCode: response.RCSuccess,
		Description:  response.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         data,
	}
	d.appLog.LogResponse(http.StatusOK, ctx.Request, req, resp)
	trx.Context.SetCustom(response.APMTrxCustomKeyResponse, resp)
	response.SendSuccessResponse(ctx, resp)
}

func (d *Delivery) SavePhoto(ctx *gin.Context) {
	span, apmCtx := apm.StartSpan(ctx.Request.Context(), "SavePhoto", response.APMSpanTypeHandlerProcess)
	defer span.End()
	trx := apm.TransactionFromContext(apmCtx)

	var req SavePhoto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		d.appLog.LogResponse(http.StatusBadRequest, ctx.Request, req, response.ErrBadRequest)
		trx.Context.SetCustom(response.APMTrxCustomKeyResponse, response.ErrBadRequest)
		response.SendErrorResponse4xx(ctx, http.StatusBadRequest, response.ErrBadRequest)
		return
	}

	d.appLog.LogRequest(ctx.Request, req)
	trx.Context.SetCustom(response.APMTrxCustomKeyRequestParam, req)

	data, err := d.service.PostingPhoto(ctx, req)
	if err != nil {
		switch err.Error() {
		case ErrUnAuthorized.Error():
			d.appLog.LogResponse(http.StatusUnauthorized, ctx.Request, req, ErrUnAuthorized)
			trx.Context.SetCustom(response.APMTrxCustomKeyResponse, ErrUnAuthorized)
			response.SendErrorResponse4xx(ctx, http.StatusUnauthorized, ErrUnAuthorized)
			return
		default:
			resp := response.ErrorStruct{
				Code:        response.RCErrorGeneral,
				Description: response.DescriptionFailed,
				Message:     err.Error(),
			}
			d.appLog.LogResponse(http.StatusInternalServerError, ctx.Request, req, resp)
			trx.Context.SetCustom(response.APMTrxCustomKeyResponse, resp)
			response.SendErrorResponse5xx(ctx, resp)
			return
		}
	}

	resp := response.Response{
		ResponseCode: response.RCSuccess,
		Description:  response.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         data,
	}
	d.appLog.LogResponse(http.StatusOK, ctx.Request, req, resp)
	trx.Context.SetCustom(response.APMTrxCustomKeyResponse, resp)
	response.SendSuccessResponse(ctx, resp)
}

func (d *Delivery) EditPhoto(ctx *gin.Context) {
	span, apmCtx := apm.StartSpan(ctx.Request.Context(), "EditPhoto", response.APMSpanTypeHandlerProcess)
	defer span.End()
	trx := apm.TransactionFromContext(apmCtx)

	var req EditPhoto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		d.appLog.LogResponse(http.StatusBadRequest, ctx.Request, req, response.ErrBadRequest)
		trx.Context.SetCustom(response.APMTrxCustomKeyResponse, response.ErrBadRequest)
		response.SendErrorResponse4xx(ctx, http.StatusBadRequest, response.ErrBadRequest)
		return
	}

	d.appLog.LogRequest(ctx.Request, req)
	trx.Context.SetCustom(response.APMTrxCustomKeyRequestParam, req)

	data, err := d.service.EditPhoto(ctx, req)
	if err != nil {
		switch err.Error() {
		case ErrUnAuthorized.Error():
			d.appLog.LogResponse(http.StatusUnauthorized, ctx.Request, req, ErrUnAuthorized)
			trx.Context.SetCustom(response.APMTrxCustomKeyResponse, ErrUnAuthorized)
			response.SendErrorResponse4xx(ctx, http.StatusUnauthorized, ErrUnAuthorized)
			return
		default:
			resp := response.ErrorStruct{
				Code:        response.RCErrorGeneral,
				Description: response.DescriptionFailed,
				Message:     err.Error(),
			}
			d.appLog.LogResponse(http.StatusInternalServerError, ctx.Request, req, resp)
			trx.Context.SetCustom(response.APMTrxCustomKeyResponse, resp)
			response.SendErrorResponse5xx(ctx, resp)
			return
		}
	}

	resp := response.Response{
		ResponseCode: response.RCSuccess,
		Description:  response.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         data,
	}
	d.appLog.LogResponse(http.StatusOK, ctx.Request, req, resp)
	trx.Context.SetCustom(response.APMTrxCustomKeyResponse, resp)
	response.SendSuccessResponse(ctx, resp)
}
