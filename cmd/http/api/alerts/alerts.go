package alerts

import (
	"context"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alerts"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/alerts"
	"git.zqbjj.top/pet/services/cmd/http/utils"
)

// GetAlertList .
// @router /api/devices/alerts [GET]
func GetAlertList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req alerts.AlertsFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetAlertListService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetAlertDetail .
// @router /api/devices/alerts/detail [GET]
func GetAlertDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewGetAlertDetailService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateAlertInfo .
// @router /api/devices/alerts [PUT]
func UpdateAlertInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req alerts.AlertInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUpdateAlertInfoService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UploadAlertInfo .
// @router /api/devices/alerts [POST]
func UploadAlertInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req alerts.AlertInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUploadAlertInfoService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteAlert .
// @router /api/devices/alerts [DELETE]
func DeleteAlert(ctx context.Context, c *app.RequestContext) {
	var err error
	var req req.IdReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewDeleteAlertService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
