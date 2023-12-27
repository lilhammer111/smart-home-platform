package alert

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/alert/handler"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"net/http"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetAlertList .
// @id			GetAlertList
// @Summary		get alert list
// @Tags		alerts
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		page	query	string	false	"page"
// @Param		limit	query	string	false	"limit"
// @Param		device_id query string false    "device id"
// @Param		level	query	string	false	"level"
// @Param		sorts	query	string	false	"sorts"
// @Param		start_date	query	string	false	"start date"
// @Param		end_date	query	string	false	"end date"
// @Success		200				{object}		example.RespOk{data=example.AlertData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/alerts/list [GET]
func GetAlertList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req alert.AlertFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetAlertListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetAlertDetail .
// @id			GetAlertDetail
// @Summary		get alert detail
// @Tags		alerts
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id		query	int	true	"id"
// @Success		200				{object}		example.RespOk{data=example.AlertData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/alerts/detail [GET]
func GetAlertDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetAlertDetailService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateAlertInfo .
// @id			UpdateAlertInfo
// @Summary		update alert info
// @Tags		alerts
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		users	body	example.AlertData	true	"alert data"
// @Success		200				{object}		example.RespOk{data=example.AlertData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/alerts/update [PUT]
func UpdateAlertInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req alert.AlertInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateAlertInfoService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UploadAlertInfo .
// @id			UploadAlertInfo
// @Summary		upload alert info
// @Tags		alerts
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		users	body	example.AlertData	true	"alert data"
// @Success		200				{object}		example.RespOk{data=example.AlertData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/alerts/upload [POST]
func UploadAlertInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req alert.AlertInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUploadAlertInfoService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteAlert .
// @id			DeleteAlert
// @Summary		delete alert info
// @Tags		alerts
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"alert id"
// @Success		200				{object}		example.RespOk 						"success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/devices/alerts/delete [DELETE]
func DeleteAlert(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteAlertService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
